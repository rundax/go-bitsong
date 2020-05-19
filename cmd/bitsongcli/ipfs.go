package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bitsongofficial/go-bitsong/ipfs"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/user"
)

func ipfsCmd() *cobra.Command {
	ipfsCmd := &cobra.Command{
		Use:                        "ipfs",
		Aliases:                    []string{"i"},
		Short:                      "Manage ipfs operations",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	ipfsCmd.AddCommand(
		ipfsInitCmd(),
		flags.LineBreak,
		ipfsPutCmd(),
		ipfsGetCmd(),
		ipfsHasCmd(),
	)

	return ipfsCmd
}

func getInstance() (*ipfs.Peer, func() error, error) {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	iden, _ := ipfs.LoadIdentity(usr.HomeDir + "/.bitsongd/ipfs/identity.json")
	if iden == nil {
		return nil, nil, fmt.Errorf("IPFS identity not found.")
	}

	path := usr.HomeDir + "/.bitsongd/ipfs/datastore"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	ds, err := ipfs.BadgerDatastore(path)
	if err != nil {
		return nil, nil, err
	}

	// listen only localhost
	listen, _ := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/4005")

	ctx := context.Background()
	h, dht, err := ipfs.SetupLibp2p(
		ctx,
		iden.PrivKey,
		nil,
		[]multiaddr.Multiaddr{listen},
		nil,
		ipfs.Libp2pOptionsExtra...,
	)

	if err != nil {
		return nil, nil, err
	}

	lite, err := ipfs.New(ctx, ds, h, dht, nil)
	if err != nil {
		return nil, nil, err
	}

	close := func() error {
		ctx.Done()
		if err := dht.Close(); err != nil {
			return err
		}
		if err := h.Close(); err != nil {
			return err
		}
		return ds.Close()
	}

	// do not bootstrap
	//lite.Bootstrap(ipfs.DefaultBootstrapPeers())

	return lite, close, nil
}

func ipfsInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize IPFS on local computer",
		RunE: func(cmd *cobra.Command, args []string) error {

			usr, err := user.Current()
			if err != nil {
				panic(err)
			}

			path := usr.HomeDir + "/.bitsongd/ipfs"
			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.MkdirAll(path, os.ModePerm)
			}

			iden, _ := ipfs.LoadIdentity(path + "/identity.json")
			if iden == nil {
				iden = &ipfs.Identity{}

				iden.PrivKey, iden.PubKey, err = crypto.GenerateKeyPair(crypto.Ed25519, 0)
				if err != nil {
					return err
				}

				err = iden.Save(path + "/identity.json")
				if err != nil {
					return err
				}

				bzPubKey, err := crypto.MarshalPublicKey(iden.PubKey)
				if err != nil {
					return err
				}
				pubkey, err := json.Marshal(bzPubKey)
				cmd.Println(fmt.Sprintf("IPFS Key generated %s", pubkey))
				return nil
			}

			bzPubKey, err := crypto.MarshalPublicKey(iden.PubKey)
			if err != nil {
				return err
			}
			pubkey, err := json.Marshal(bzPubKey)
			cmd.Println(fmt.Sprintf("IPFS Key is already generated %s", pubkey))
			return nil
		},
	}

	return cmd
}

func ipfsPutCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put [file]",
		Short: "Put a new file",
		Long:  fmt.Sprintf(`%s`, version.ClientName),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			peer, close, err := getInstance()
			file, _ := os.Open(args[0])
			n, err := peer.AddFile(context.Background(), bufio.NewReader(file), nil)
			if err != nil {
				return err
			}

			for _, l := range n.Links() {
				fmt.Println(fmt.Sprintf("cid: %s size: %d name: %s", l.Cid, l.Size, l.Name))
			}

			fmt.Println(n.Cid())

			if err = close(); err != nil {
				return fmt.Errorf("error when closing the client: %v", err)
			}

			return nil
		},
	}

	return cmd
}

func ipfsGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [cid]",
		Short: "Get cid from datastore",
		Long:  fmt.Sprintf(`%s`, version.ClientName),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			peer, close, err := getInstance()
			if err != nil {
				return err
			}

			c, err := cid.Decode(args[0])
			if err != nil {
				return err
			}

			ctx := context.Background()
			rsc, err := peer.GetFile(ctx, c)
			if err != nil {
				panic(err)
			}
			defer rsc.Close()
			content, err := ioutil.ReadAll(rsc)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(content))

			if err = close(); err != nil {
				return fmt.Errorf("error when closing the client: %v", err)
			}

			return nil
		},
	}

	return cmd
}

func ipfsHasCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "has [cid]",
		Short: "Check if cid is stored in local datastore",
		Long:  fmt.Sprintf(`%s`, version.ClientName),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			peer, close, err := getInstance()
			if err != nil {
				return err
			}

			c, err := cid.Decode(args[0])
			if err != nil {
				return err
			}

			found, err := peer.HasBlock(c)
			if err != nil {
				return err
			}
			fmt.Println(found)

			if err = close(); err != nil {
				return fmt.Errorf("error when closing the client: %v", err)
			}

			return nil
		},
	}

	return cmd
}

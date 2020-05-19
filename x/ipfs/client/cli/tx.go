package cli

import (
	"bufio"
	"context"
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/ipfs/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkctx "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
	"github.com/spf13/cobra"
	"os"
	"os/user"

	ipfs "github.com/bitsongofficial/go-bitsong/types/ipfs"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	trackTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	trackTxCmd.AddCommand(flags.PostCommands(
		GetCmdInit(cdc),
		GetCmdPut(cdc),
		GetCmdGet(cdc),
	)...)

	return trackTxCmd
}

func setupLite(priv crypto.PrivKey) (*ipfs.Peer, func() error, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, nil, err
	}

	path := usr.HomeDir + "/.bitsongd/ipfs/datastore"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	ds, err := ipfs.BadgerDatastore(path)
	if err != nil {
		return nil, nil, err
	}

	listen, _ := multiaddr.NewMultiaddr("/ip4/0.0.0.0/tcp/4005")

	ctx := context.Background()
	h, dht, err := ipfs.SetupLibp2p(
		ctx,
		priv,
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

	lite.Bootstrap(ipfs.DefaultBootstrapPeers())

	return lite, close, nil
}

func GetCmdInit(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Init ipfs and identity",
		Long:  fmt.Sprintf(`%s`, version.ClientName),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			_ = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			_ = sdkctx.NewCLIContextWithInput(inBuf).WithCodec(cdc)

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

				fmt.Println(fmt.Sprintf(`IPFS Key generated`))
				return nil
			}

			fmt.Println(fmt.Sprintf(`IPFS Key is already generated`))
			return nil
		},
	}

	return cmd
}

func GetCmdPut(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put [file]",
		Short: "Put a new file",
		Long:  fmt.Sprintf(`%s`, version.ClientName),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			_ = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			_ = sdkctx.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			usr, err := user.Current()
			if err != nil {
				panic(err)
			}

			iden, _ := ipfs.LoadIdentity(usr.HomeDir + "/.bitsongd/ipfs/identity.json")
			if iden == nil {
				return fmt.Errorf("IPFS identity not found.")
			}

			lite, close, err := setupLite(iden.PrivKey)
			file, _ := os.Open(args[0])
			n, err := lite.AddFile(context.Background(), bufio.NewReader(file), nil)
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

func GetCmdGet(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [cid]",
		Short: "Get cid from datastore",
		Long:  fmt.Sprintf(`%s`, version.ClientName),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			_ = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			_ = sdkctx.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			usr, err := user.Current()
			if err != nil {
				panic(err)
			}

			iden, _ := ipfs.LoadIdentity(usr.HomeDir + "/.bitsongd/ipfs/identity.json")
			if iden == nil {
				return fmt.Errorf("IPFS identity not found.")
			}

			lite, close, err := setupLite(iden.PrivKey)
			if err != nil {
				return err
			}

			c, err := cid.Decode(args[0])
			if err != nil {
				return err
			}

			found, err := lite.HasBlock(c)
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

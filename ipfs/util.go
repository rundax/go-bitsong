package ipfs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/multiformats/go-multiaddr"
	"io/ioutil"
	"os"
	"os/user"
	"time"

	"github.com/ipfs/go-datastore"
	badger "github.com/ipfs/go-ds-badger"
	config "github.com/ipfs/go-ipfs-config"
	ipns "github.com/ipfs/go-ipns"
	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	host "github.com/libp2p/go-libp2p-core/host"
	peer "github.com/libp2p/go-libp2p-core/peer"
	pnet "github.com/libp2p/go-libp2p-core/pnet"
	routing "github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	dualdht "github.com/libp2p/go-libp2p-kad-dht/dual"
	record "github.com/libp2p/go-libp2p-record"
	secio "github.com/libp2p/go-libp2p-secio"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
)

// DefaultBootstrapPeers returns the default go-ipfs bootstrap peers (for use
// with NewLibp2pHost.
func DefaultBootstrapPeers() []peer.AddrInfo {
	defaults, _ := config.DefaultBootstrapPeers()
	return defaults
}

// IPFSBadgerDatastore returns the Badger datastore used by the IPFS daemon
// (from `~/.ipfs/datastore`). Do not use the default datastore when the
// regular IFPS daemon is running at the same time.
func IPFSBadgerDatastore() (datastore.Batching, error) {
	home := os.Getenv("HOME")
	if home == "" {
		usr, err := user.Current()
		if err != nil {
			panic(fmt.Sprintf("cannot get current user: %s", err))
		}
		home = usr.HomeDir
	}

	path, err := config.DataStorePath(home)
	if err != nil {
		return nil, err
	}
	return BadgerDatastore(path)
}

// BadgerDatastore returns a new instance of Badger-DS persisting
// to the given path with the default options.
func BadgerDatastore(path string) (datastore.Batching, error) {
	return badger.NewDatastore(path, &badger.DefaultOptions)
}

// Libp2pOptionsExtra provides some useful libp2p options
// to create a fully featured libp2p host. It can be used with
// SetupLibp2p.
var Libp2pOptionsExtra = []libp2p.Option{
	libp2p.NATPortMap(),
	libp2p.ConnectionManager(connmgr.NewConnManager(100, 600, time.Minute)),
	libp2p.EnableAutoRelay(),
	libp2p.EnableNATService(),
	libp2p.Security(libp2ptls.ID, libp2ptls.New),
	libp2p.Security(secio.ID, secio.New),
	// TODO: re-enable when QUIC support private networks.
	// libp2p.Transport(libp2pquic.NewTransport),
	libp2p.DefaultTransports,
}

// SetupLibp2p returns a routed host and DHT instances that can be used to
// easily create a ipfslite Peer. You may consider to use Peer.Bootstrap()
// after creating the IPFS-Lite Peer to connect to other peers. When the
// datastore parameter is nil, the DHT will use an in-memory datastore, so all
// provider records are lost on program shutdown.
//
// Additional libp2p options can be passed. Note that the Identity,
// ListenAddrs and PrivateNetwork options will be setup automatically.
// Interesting options to pass: NATPortMap() EnableAutoRelay(),
// libp2p.EnableNATService(), DisableRelay(), ConnectionManager(...)... see
// https://godoc.org/github.com/libp2p/go-libp2p#Option for more info.
//
// The secret should be a 32-byte pre-shared-key byte slice.
func SetupLibp2p(
	ctx context.Context,
	hostKey crypto.PrivKey,
	secret pnet.PSK,
	listenAddrs []multiaddr.Multiaddr,
	ds datastore.Batching,
	opts ...libp2p.Option,
) (host.Host, *dualdht.DHT, error) {

	var ddht *dualdht.DHT
	var err error

	finalOpts := []libp2p.Option{
		libp2p.Identity(hostKey),
		libp2p.ListenAddrs(listenAddrs...),
		libp2p.PrivateNetwork(secret),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			ddht, err = newDHT(ctx, h, ds)
			return ddht, err
		}),
	}
	finalOpts = append(finalOpts, opts...)

	h, err := libp2p.New(
		ctx,
		finalOpts...,
	)
	if err != nil {
		return nil, nil, err
	}

	return h, ddht, nil
}

type Identity struct {
	PrivKey crypto.PrivKey
	PubKey  crypto.PubKey
}

type jsonIdentity struct {
	PrivKey, PubKey []byte
}

// LoadIdentity loads a json config from a path.
func LoadIdentity(path string) (*Identity, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	jsonID := &jsonIdentity{}
	err = json.Unmarshal(file, jsonID)
	if err != nil {
		return nil, err
	}

	c := &Identity{}
	c.PrivKey, err = crypto.UnmarshalPrivateKey(jsonID.PrivKey)
	if err != nil {
		return nil, err
	}
	c.PubKey, err = crypto.UnmarshalPublicKey(jsonID.PubKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return c, nil
}

func (c *Identity) Save(path string) error {
	if c == nil {
		return fmt.Errorf("Identity is nil")
	}
	if c.PrivKey == nil {
		return fmt.Errorf("PrivKey is nil")
	}
	bzPrivKey, err := crypto.MarshalPrivateKey(c.PrivKey)
	if err != nil {
		return err
	}
	bzPubKey, err := crypto.MarshalPublicKey(c.PubKey)
	if err != nil {
		return err
	}
	jsonID := &jsonIdentity{PrivKey: bzPrivKey, PubKey: bzPubKey}
	file, err := json.MarshalIndent(jsonID, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, file, 0644)
}

func newDHT(ctx context.Context, h host.Host, ds datastore.Batching) (*dualdht.DHT, error) {
	dhtOpts := []dht.Option{
		dht.NamespacedValidator("pk", record.PublicKeyValidator{}),
		dht.NamespacedValidator("ipns", ipns.Validator{KeyBook: h.Peerstore()}),
		dht.Concurrency(10),
		dht.Mode(dht.ModeAuto),
	}
	if ds != nil {
		dhtOpts = append(dhtOpts, dht.Datastore(ds))
	}

	return dualdht.New(ctx, h, dhtOpts...)

}

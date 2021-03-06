package mockrouting

import (
	mocknet "github.com/heems/bssim/Godeps/_workspace/src/github.com/ipfs/go-ipfs/p2p/net/mock"
	dht "github.com/heems/bssim/Godeps/_workspace/src/github.com/ipfs/go-ipfs/routing/dht"
	"github.com/heems/bssim/Godeps/_workspace/src/github.com/ipfs/go-ipfs/util/testutil"
	ds "github.com/heems/bssim/Godeps/_workspace/src/github.com/jbenet/go-datastore"
	sync "github.com/heems/bssim/Godeps/_workspace/src/github.com/jbenet/go-datastore/sync"
	context "github.com/heems/bssim/Godeps/_workspace/src/golang.org/x/net/context"
)

type mocknetserver struct {
	mn mocknet.Mocknet
}

func NewDHTNetwork(mn mocknet.Mocknet) Server {
	return &mocknetserver{
		mn: mn,
	}
}

func (rs *mocknetserver) Client(p testutil.Identity) Client {
	return rs.ClientWithDatastore(context.TODO(), p, ds.NewMapDatastore())
}

func (rs *mocknetserver) ClientWithDatastore(ctx context.Context, p testutil.Identity, ds ds.Datastore) Client {

	// FIXME AddPeer doesn't appear to be idempotent

	host, err := rs.mn.AddPeer(p.PrivateKey(), p.Address())
	if err != nil {
		panic("FIXME")
	}
	return dht.NewDHT(ctx, host, sync.MutexWrap(ds))
}

var _ Server = &mocknetserver{}

package node

import (
	"bytes"
	"context"
	"errors"
	"github.com/dgraph-io/badger/v3"
	"github.com/itsabgr/statelog/internal/pb"
	"github.com/itsabgr/statelog/pkg/api"
	"github.com/itsabgr/statelog/pkg/crypto"
	"github.com/itsabgr/statelog/pkg/event"
	"github.com/itsabgr/statelog/pkg/storage"
	"github.com/itsabgr/statelog/pkg/tmp"
	"github.com/itsabgr/xorcontacts"
)

type Node struct {
	db    *badger.DB
	tmp   *tmp.Tmp
	sk    []byte
	cli   *api.Client
	peers xorcontacts.List[*Peer]
}

func (n *Node) PK() []byte {
	return crypto.DerivePK(n.sk)
}
func (n *Node) Sign(data []byte) []byte {
	return crypto.Sign(n.sk, data)
}
func (n *Node) Close() error {
	return n.db.Close()
}
func (n *Node) validateEvent(ev *event.Event) {

}
func (n *Node) Publish(ctx context.Context, ev *event.Event) error {
	peers := n.peers.Xor(&Peer{pk: ev.Signer})
	var errorList []error
	for _, peer := range peers {
		out, err := n.cli.PushEvent(peer.addr, ctx, &pb.PushEventInput{Event: (*pb.Event)(ev)})
		if err != nil {
			errorList = append(errorList, err)
			continue
		}
		if !bytes.Equal(out.Signer, peer.pk) {
			errorList = append(errorList, errors.New("invalid signer"))
			continue
		}
		err = crypto.Verify(out.Signer, ev.Encode(), out.Sig)
		if err != nil {
			errorList = append(errorList, err)
			continue
		}
	}
	if len(errorList)/n.peers.Len() < 1/3 {
		tx := n.db.NewTransaction(true)
		defer tx.Discard()
		err := storage.Wrap(tx).InsertEvent(ev)
		if err != nil {
			return err
		}
		err = tx.Commit()
		if err != nil {
			return err
		}
	}
	n.tmp.Delete(ev.Signer, ev.Index)
	return nil
}
func (n *Node) Push(ctx context.Context, event *event.Event) error {
	err := event.VerifySig()
	if err != nil {
		return err
	}
	ev, _, err := n.Get(ctx, event.Signer, event.Index)
	if err == nil && ev != nil {
		return nil
	}
	//

	//
	return nil
}
func (n *Node) Get(ctx context.Context, signer []byte, index uint64) (ev *event.Event, confirm bool, err error) {
	ev = n.tmp.Get(signer, index)
	if ev != nil {
		return ev, false, nil
	}
	if ctx.Err() != nil {
		return nil, false, ctx.Err()
	}
	tx := n.db.NewTransaction(false)
	defer tx.Discard()
	ev, err = storage.Wrap(tx).GetEvent(signer, index)
	return ev, true, err
}

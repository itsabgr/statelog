package event

import (
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/itsabgr/go-handy"
	"github.com/itsabgr/statelog/internal/pb"
	"github.com/itsabgr/statelog/pkg/crypto"
)

type Event pb.Event

type Ptr pb.EventPtr

func (ptr *Ptr) ID() []byte {
	return EncodeID(ptr.Signer, ptr.Index)
}
func EncodeID(signer []byte, index uint64) []byte {
	b := make([]byte, len(signer)+8)
	binary.BigEndian.PutUint64(b[len(signer):], index)
	copy(b[len(signer):], signer)
	return b
}

var ErrInvalidID = errors.New("invalid id")

func DecodeID(b []byte) (signer []byte, index uint64, err error) {
	if len(b) <= 8 {
		return nil, 0, ErrInvalidID
	}
	index = binary.BigEndian.Uint64(b[len(b)-8:])
	signer = signer[:len(b)-8]
	return
}

func DecodePtr(b []byte) (*Ptr, error) {
	ev := &pb.EventPtr{}
	return (*Ptr)(ev), proto.Unmarshal(b, ev)
}

func DecodeEvent(b []byte) (*Event, error) {
	ev := &pb.Event{}
	return (*Event)(ev), proto.Unmarshal(b, ev)
}
func (ev *Event) ID() []byte {
	return EncodeID(ev.Signer, ev.Index)
}
func (ev *Event) Ptr() *Ptr {
	ptr := &Ptr{}
	ptr.Signer = ev.Signer
	ptr.Index = ev.Index
	ptr.Digest = ev.Digest()
	return ptr
}
func (ev *Event) Encode() []byte {
	out, err := proto.Marshal((*pb.Event)(ev))
	handy.Throw(err)
	return out
}
func (ev *Event) Digest() []byte {
	return crypto.Hash(ev.Encode())
}
func (ev *Event) VerifySig() error {
	sig := ev.Sig
	ev.Sig = nil
	err := crypto.Verify(ev.Signer, ev.Encode(), sig)
	ev.Sig = sig
	return err
}

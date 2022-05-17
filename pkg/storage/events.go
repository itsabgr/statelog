package storage

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/dgraph-io/badger/v3"
	"github.com/itsabgr/go-handy"
	"github.com/itsabgr/statelog/pkg/event"
	"github.com/itsabgr/statelog/pkg/utils"
)

const (
	tableAccounts byte = 5
	tablePeers    byte = 4
	tableEvents   byte = 3
	tablePtrs     byte = 2
	tableFins     byte = 1
)

type tx badger.Txn

func (t *tx) unwrap() *badger.Txn {
	return (*badger.Txn)(t)
}

func Wrap(t *badger.Txn) *tx {
	return (*tx)(t)
}

func (t *tx) setEvent(ev *event.Event) error {
	lastIndex, err := t.GetLastEventIndex(ev.Signer)
	if err != nil {
		return err
	}
	err = t.set(tableEvents, ev.ID(), ev.Encode())
	if err != nil {
		return err
	}
	if lastIndex == 0 || lastIndex >= ev.Index {
		return nil
	}
	//set last event index
	return t.set(tableEvents, event.EncodeID(ev.Signer, 0), utils.Uint64ToBE(ev.Index))
}
func (t *tx) GetFirstEvent(singer []byte) (*event.Event, error) {
	return t.GetNextEvent(singer, 0, false)
}
func (t *tx) GetLastEvent(signer []byte) (*event.Event, error) {
	lastIndex, err := t.GetLastEventIndex(signer)
	if err != nil {
		return nil, err
	}
	return t.GetEvent(signer, lastIndex)
}
func (t *tx) GetLastEventIndex(signer []byte) (uint64, error) {
	lastIndexB, err := t.get(tableEvents, event.EncodeID(signer, 0))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return 0, nil
		}
		return 0, err
	}
	lastIndex := binary.BigEndian.Uint64(lastIndexB)
	return lastIndex, nil
}

func (t *tx) InsertEvent(ev *event.Event) error {
	err := ev.VerifySig()
	if err != nil {
		return err
	}
	for _, ptr := range ev.Events {
		ev2, err := t.GetEvent(ptr.Signer, ptr.Index)
		if err != nil {
			return errors.New("missing ptr")
		}
		if ev.Index <= ev2.Index {
			return errors.New("conflict time")
		}
		if false == bytes.Equal(ev2.Digest(), ptr.Digest) {
			return errors.New("invalid ptr hash")
		}
	}
	//
	has, err := t.HasEvent(ev.Signer, ev.Index)
	if err != nil {
		return err
	}
	if has {
		return badger.ErrConflict
	}
	err = t.setEvent(ev)
	if err != nil {
		return err
	}
	for _, ptr := range ev.Events {
		_ = t.addPtr(ev.Signer, ev.Index, ptr.Signer, ptr.Index)
	}
	return nil
}

func (t *tx) get(table byte, key []byte) ([]byte, error) {
	item, err := t.unwrap().Get(append([]byte{table}, key...))
	if err != nil {
		return nil, err
	}
	if item.IsDeletedOrExpired() {
		return nil, badger.ErrKeyNotFound
	}
	val, err := item.ValueCopy(nil)
	handy.Throw(err)
	return val, nil
}
func (t *tx) has(table byte, key []byte) (bool, error) {
	item, err := t.unwrap().Get(append([]byte{table}, key...))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return false, nil
		}
		return false, err
	}
	if item.IsDeletedOrExpired() {
		return false, nil
	}
	return true, nil
}
func (t *tx) set(table byte, key, val []byte) error {
	return t.unwrap().Set(append([]byte{table}, key...), val)

}
func (t *tx) HasEvent(singer []byte, index uint64) (bool, error) {
	return t.has(tableEvents, event.EncodeID(singer, index))
}

func (t *tx) GetEvent(singer []byte, index uint64) (*event.Event, error) {
	val, err := t.get(tableEvents, event.EncodeID(singer, index))
	if err != nil {
		return nil, err
	}
	ev, err := event.DecodeEvent(val)
	handy.Throw(err)
	return ev, nil
}

func (t *tx) GetNextEvent(singer []byte, index uint64, rev bool) (*event.Event, error) {
	val, err := t.getNext(tableEvents, event.EncodeID(singer, index), rev)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}
	ev, err := event.DecodeEvent(val)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(singer, ev.Signer) {
		return nil, nil
	}
	return ev, nil
}

func (t *tx) getNext(table byte, key []byte, rev bool) ([]byte, error) {
	itr := t.unwrap().NewIterator(badger.IteratorOptions{Prefix: []byte{table}, Reverse: rev})
	defer itr.Close()
	for itr.Seek(key); itr.Valid(); itr.Next() {
		val, err := itr.Item().ValueCopy(nil)
		handy.Throw(err)
		if !bytes.Equal(val, key) {
			return val, nil
		}
	}
	return nil, nil
}

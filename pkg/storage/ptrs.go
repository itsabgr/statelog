package storage

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/itsabgr/go-handy"
	"github.com/itsabgr/statelog/pkg/event"
	"github.com/itsabgr/statelog/pkg/utils"
)

func (t *tx) addPtrR(fromSig []byte, fromIndex uint64, toSig []byte, toIndex uint64) (wasFin bool, _ error) {
	fin, err := t.IsFin(toSig, toIndex)
	if err != nil {
		return false, err
	}
	if fin {
		return true, nil
	}
	has, err := t.has(tablePtrs, handy.Concat(event.EncodeID(toSig, toIndex), fromSig))
	if err != nil {
		return false, err
	}
	if has {
		return false, nil
	}
	return false, t.set(tablePtrs, handy.Concat(event.EncodeID(toSig, toIndex), fromSig), utils.Uint64ToBE(fromIndex))
}

func (t *tx) addPtr(fromSig []byte, fromIndex uint64, toSig []byte, toIndex uint64) error {
	wasFin, err := t.addPtrR(fromSig, fromIndex, toSig, toIndex)
	if err != nil {
		return err
	}
	if wasFin {
		return nil
	}
	ev, err := t.GetEvent(toSig, toIndex)
	if err != nil {
		return err
	}
	for _, ptr := range ev.Events {
		err := t.addPtr(fromSig, fromIndex, ptr.Signer, ptr.Index)
		if err != nil {
			return err
		}
	}
	return nil
}
func (t *tx) countPrefix(table byte, prefix []byte) uint64 {
	itr := t.unwrap().NewIterator(badger.IteratorOptions{Prefix: []byte{table}})
	defer itr.Close()
	count := uint64(0)
	for ; itr.ValidForPrefix(prefix); itr.Next() {
		count++
	}
	return count
}
func (t *tx) CountPtrs(signer []byte, index uint64) uint64 {
	return t.countPrefix(tablePtrs, event.EncodeID(signer, index))
}

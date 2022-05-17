package storage

import (
	"encoding/binary"
	"github.com/dgraph-io/badger/v3"
	"github.com/itsabgr/statelog/pkg/event"
	"github.com/itsabgr/statelog/pkg/utils"
	"time"
)

func (t *tx) setFinI(signer []byte, index uint64, at time.Time) (bool, error) {
	has, err := t.has(tableFins, event.EncodeID(signer, index))
	if err != nil {
		return false, err
	}
	if has {
		return true, nil
	}
	lastFinIndex, err := t.GetLastFinIndex(signer)
	if err != nil {
		return false, err
	}
	err = t.set(tableFins, event.EncodeID(signer, index), utils.Uint64ToBE(uint64(at.UnixNano())))
	if err != nil {
		return false, err
	}
	if lastFinIndex < index {
		return false, t.set(tableFins, event.EncodeID(signer, 0), utils.Uint64ToBE(index))
	}
	return false, nil
}
func (t *tx) SetFin(signer []byte, index uint64, at time.Time) error {
	was, err := t.setFinI(signer, index, at)
	if err != nil {
		return err
	}
	if was {
		return nil
	}
	ev, err := t.GetEvent(signer, index)
	if err != nil {
		return err
	}
	for _, ptr := range ev.Events {
		err := t.SetFin(ptr.Signer, ptr.Index, at)
		if err != nil {
			return err
		}
	}
	return nil
}
func (t *tx) IsFin(signer []byte, index uint64) (bool, error) {
	has, err := t.has(tableFins, event.EncodeID(signer, index))
	if err != nil {
		return false, err
	}
	return has, err
}
func (t *tx) GetLastFinIndex(signer []byte) (uint64, error) {
	lastFinIndex, err := t.get(tableFins, event.EncodeID(signer, 0))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return 0, nil
		}
		return 0, err
	}
	return binary.BigEndian.Uint64(lastFinIndex), nil
}
func (t *tx) GetLastFin(signer []byte) (*event.Event, error) {
	index, err := t.GetLastFinIndex(signer)
	if err != nil {
		return nil, err
	}
	return t.GetEvent(signer, index)
}

package storage

import (
	"encoding/binary"
	"github.com/dgraph-io/badger/v3"
	"github.com/itsabgr/go-handy"
	"github.com/itsabgr/statelog/pkg/uid"
	"github.com/itsabgr/statelog/pkg/utils"
)

func (t *tx) AddAccount(pk []byte, delta int64) error {
	return t.set(tableAccounts, handy.Concat(pk, utils.Uint64ToBE(uint64(uid.Gen()))), utils.Uint64ToBE(uint64(delta)))
}
func (t *tx) SumAccounts(pk []byte) (int64, error) {
	accounts, err := t.list2(tableAccounts, pk)
	if err != nil {
		return 0, err
	}
	sum := int64(0)
	for _, account := range accounts {
		sum += int64(binary.BigEndian.Uint64(account))
	}
	return sum, nil
}
func (t *tx) VacAccounts(pk []byte) error {
	itr := t.unwrap().NewIterator(badger.IteratorOptions{Prefix: append([]byte{tableAccounts}, pk...)})
	defer itr.Close()
	var sum int64
	for itr.Rewind(); itr.Valid(); itr.Next() {
		val, err := itr.Item().ValueCopy(nil)
		if err != nil {
			return err
		}
		sum += int64(binary.BigEndian.Uint64(val))
		err = t.delete(tableAccounts, itr.Item().Key())
		if err != nil {
			return err
		}
	}
	return t.AddAccount(pk, sum)
}

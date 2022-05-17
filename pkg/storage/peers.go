package storage

import "github.com/dgraph-io/badger/v3"

func (t *tx) AddPeer(pk []byte, data []byte) error {
	return t.set(tablePeers, pk, data)
}
func (t *tx) DeletePeer(pk []byte) error {
	return t.delete(tablePeers, pk)
}
func (t *tx) delete(table byte, key []byte) error {
	return t.unwrap().Delete(append([]byte{table}, key...))
}
func (t *tx) list(table byte) ([][]byte, error) {
	itr := t.unwrap().NewIterator(badger.IteratorOptions{Prefix: []byte{table}})
	defer itr.Close()
	var b [][]byte
	for itr.Rewind(); itr.Valid(); itr.Next() {
		val, err := itr.Item().ValueCopy(nil)
		if err != nil {
			return nil, err
		}
		b = append(b, val)
	}
	return b, nil
}
func (t *tx) list2(table byte, prefix []byte) ([][]byte, error) {
	itr := t.unwrap().NewIterator(badger.IteratorOptions{Prefix: append([]byte{table}, prefix...)})
	defer itr.Close()
	var b [][]byte
	for itr.Rewind(); itr.Valid(); itr.Next() {
		val, err := itr.Item().ValueCopy(nil)
		if err != nil {
			return nil, err
		}
		b = append(b, val)
	}
	return b, nil
}
func (t *tx) ListPeers() ([][]byte, error) {
	return t.list(tablePeers)
}

package tmp

import (
	"encoding/gob"
	"encoding/hex"
	"github.com/itsabgr/statelog/pkg/event"
	"io"
	"sync"
)

type Tmp struct {
	_map  map[string]*event.Event
	mutex sync.RWMutex
}

func New() *Tmp {
	return &Tmp{
		_map:  make(map[string]*event.Event),
		mutex: sync.RWMutex{},
	}
}
func (t *Tmp) Import(r io.Reader) error {
	_map := map[string]*event.Event{}
	err := gob.NewDecoder(r).Decode(_map)
	if err != nil {
		return err
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for k, v := range _map {
		t._map[k] = v
	}
	return nil
}
func (t *Tmp) Export(w io.Writer) error {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return gob.NewEncoder(w).Encode(t._map)
}
func (t *Tmp) Get(signer []byte, index uint64) *event.Event {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	ev, ok := t._map[hex.EncodeToString(event.EncodeID(signer, index))]
	if !ok {
		return nil
	}
	return ev
}
func (t *Tmp) Delete(signer []byte, index uint64) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	delete(t._map, hex.EncodeToString(event.EncodeID(signer, index)))
}

func (t *Tmp) Filter(fn func(*event.Event) bool) []*event.Event {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	b := make([]*event.Event, 0, 1)
	for _, ev := range t._map {
		if fn(ev) {
			b = append(b, ev)
		}
	}
	return b
}
func (t *Tmp) Remove(fn func(*event.Event) bool) int {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	b := make([]*event.Event, 0, 1)
	for _, ev := range t._map {
		if fn(ev) {
			b = append(b, ev)
		}
	}
	for _, ev := range b {
		delete(t._map, hex.EncodeToString(event.EncodeID(ev.Signer, ev.Index)))
	}
	return len(b)
}
func (t *Tmp) Len() int {
	return len(t._map)
}
func (t *Tmp) Count(fn func(*event.Event) bool) int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	b := 0
	for _, ev := range t._map {
		if fn(ev) {
			b += 1
		}
	}
	return b
}
func (t *Tmp) First(fn func(*event.Event) bool) *event.Event {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	for _, ev := range t._map {
		if fn(ev) {
			return ev
		}
	}
	return nil
}

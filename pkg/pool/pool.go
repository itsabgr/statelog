package pool

import (
	"sync"
	"time"
)

type ent[T any] struct {
	get time.Time
	t   T
}
type Pool[V comparable, T any] struct {
	//Incomplete
	_map  map[V]*ent[T]
	mutex sync.Mutex
	init  func(V) (T, error)
	gc    time.Duration
}

func New[V comparable, T any](init func(V) (T, error), timeout time.Duration) *Pool[V, T] {
	pool := &Pool[V, T]{
		init:  init,
		mutex: sync.Mutex{},
		_map:  make(map[V]*ent[T]),
		gc:    timeout,
	}
	return pool
}

func (p *Pool[V, T]) GC() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for k, e := range p._map {
		if e.get.Add(p.gc).Before(time.Now()) {
			delete(p._map, k)
		}
	}
}
func (p *Pool[V, T]) Get(v V) (T, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	e, ok := p._map[v]
	if ok {
		e.get = time.Now()
		return e.t, nil
	}
	t, err := p.init(v)
	if err != nil {
		return t, err
	}
	p._map[v] = &ent[T]{time.Now(), t}
	return t, nil
}

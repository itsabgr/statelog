package uid

import (
	"sync/atomic"
	"time"
)

var base = time.Now().UnixNano()

func Gen() int64 {
	return atomic.AddInt64(&base, 1)
}

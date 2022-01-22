package syncx

import (
	"sync"
)

//Shield 护盾承担保护作用
type Shield struct {
	mtx sync.Mutex
}

func (b *Shield) Guard(fn func()) {
	Guard(&b.mtx, fn)
}

func Guard(locker sync.Locker, fn func()) {
	locker.Lock()
	defer locker.Unlock()
	fn()
}

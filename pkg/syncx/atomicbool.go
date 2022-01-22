package syncx

import "sync/atomic"

// 需要的方法CompareAndSwap  Set Load

type AtomicBool uint32

func NewAtomicBool() *AtomicBool {
	return new(AtomicBool)
}

func ForAtomicBool(val bool) *AtomicBool {
	b := NewAtomicBool()
	b.Set(val)
	return b
}

func (b *AtomicBool) Set(val bool) {
	if val {
		atomic.StoreUint32((*uint32)(b), 1)
	} else {
		atomic.StoreUint32((*uint32)(b), 0)
	}
}

func (b *AtomicBool) CompareAndSwap(old, new bool) {

}

func (b *AtomicBool) IsTrue() bool {
	return atomic.LoadUint32((*uint32)(b)) == 1
}

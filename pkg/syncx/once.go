package syncx

import (
	"sync"
)

func Once(f func()) func() {
	once := new(sync.Once)
	return func() {
		once.Do(f)
	}
}

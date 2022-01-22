package syncx

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShield(t *testing.T) {
	var sum int
	fn := func() {
		sum++
	}
	s := &Shield{}
	w := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			s.Guard(fn)
		}()
	}
	w.Wait()
	assert.Equal(t, 100, sum)
}

func BenchmarkShield(b *testing.B) {
	var sum int
	fn := func() {
		sum++
	}
	s := &Shield{}
	w := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			s.Guard(fn)
		}()
	}
	w.Wait()
	assert.Equal(b, b.N, sum)
}

package syncx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type E struct {
	Sum int
}

func TestOnce(t *testing.T) {
	var sum int
	f := func() {
		sum++
	}
	fun := Once(f)
	for i := 1; i < 3; i++ {
		fun()
	}
	assert.Equal(t, 1, sum)
}

func BenchmarkOnce(b *testing.B) {
	var sum int

	fun := Once(func() {
		sum++
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fun()
	}
	assert.Equal(b, 1, sum)
}

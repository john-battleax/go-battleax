package syncx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomicBool(t *testing.T) {
	b := ForAtomicBool(true)
	assert.Equal(t, true, b.IsTrue())
	b.Set(false)
	assert.Equal(t, false, b.IsTrue())
}

func BenchmarkAtomicBool(b *testing.B) {
	a := ForAtomicBool(true)

	assert.Equal(b, true, a.IsTrue())

	for i := 0; i < b.N; i++ {
		a.Set(false)
	}
	assert.Equal(b, false, a.IsTrue())
}

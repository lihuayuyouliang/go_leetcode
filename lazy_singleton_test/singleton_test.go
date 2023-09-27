package lazy_singleton_test

import (
	"testing"

	"leetcode/lazy_singleton"

	"github.com/stretchr/testify/assert"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, lazy_singleton.GetLazyInstance(), lazy_singleton.GetLazyInstance())
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if lazy_singleton.GetLazyInstance() != lazy_singleton.GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

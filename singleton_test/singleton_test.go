package singleton_test

import (
	"github.com/stretchr/testify/assert"
	"leetcode/singleton"
	"testing"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, singleton.GetInstance(), singleton.GetInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetInstance() != singleton.GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

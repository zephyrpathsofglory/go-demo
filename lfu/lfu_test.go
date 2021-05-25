package lfu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	cache := NewCache(100)

	cache.Put(1, 1)
	val := cache.Get(1)
	assert.Equal(t, 1, val)
}

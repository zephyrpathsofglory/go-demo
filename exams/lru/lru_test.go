package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	cache := NewCache(3)

	cache.Put(1, 1)
	val := cache.Get(1)
	assert.Equal(t, 1, val)

	cache.Put(2, 2)
	val = cache.Get(2)
	assert.Equal(t, 2, val)

	cache.Put(3, 3)
	val = cache.Get(3)
	assert.Equal(t, 3, val)

	cache.Put(4, 4)
	val = cache.Get(4)
	assert.Equal(t, 4, val)

	val = cache.Get(1)
	assert.Equal(t, -1, val)

	val = cache.Get(2)
	assert.Equal(t, 2, val)

	val = cache.Get(3)
	assert.Equal(t, 3, val)

	val = cache.Get(4)
	assert.Equal(t, 4, val)

	cache.Put(5, 5)
	val = cache.Get(5)
	assert.Equal(t, 5, val)

	val = cache.Get(2)
	assert.Equal(t, -1, val)

	val = cache.Get(3)
	assert.Equal(t, 3, val)

	val = cache.Get(4)
	assert.Equal(t, 4, val)
}

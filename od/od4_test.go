package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookStack(t *testing.T) {
	assert.Equal(t, 3, bookStack([][2]int{
		{16, 15},
		{15, 14},
		{13, 12},
	}))

	assert.Equal(t, 3, bookStack([][2]int{
		{16, 15},
		{16, 14},
		{15, 14},
		{13, 12},
	}))

	assert.Equal(t, 3, bookStack([][2]int{
		{16, 15},
		{13, 12},
		{15, 14},
		{16, 14},
	}))
}

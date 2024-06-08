package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChiHuoGuo(t *testing.T) {
	assert.Equal(t, 1, chiHuoGuo([][2]int{
		{1, 2},
		{2, 1},
	}, 1))

	assert.Equal(t, 2, chiHuoGuo([][2]int{
		{1, 2},
		{2, 1},
		{3, 2},
		{4, 3},
	}, 3))

	assert.Equal(t, 3, chiHuoGuo([][2]int{
		{1, 12},
		{2, 1},
		{3, 2},
		{4, 3},
	}, 3))
}

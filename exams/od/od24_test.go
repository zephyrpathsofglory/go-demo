package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeviceWorkingTime(t *testing.T) {
	assert.Equal(t, 5, DeviceWorkingTime([][]int{
		{0, 1}, {1, 3}, {2, 2}, {3, 3}, {4, 1}, {4, 3},
	}, 3))

	assert.Equal(t, 6, DeviceWorkingTime([][]int{
		{3, 3}, {4, 1}, {5, 2}, {6, 1}, {7, 3}, {8, 1}, {11, 3},
	}, 4))
}

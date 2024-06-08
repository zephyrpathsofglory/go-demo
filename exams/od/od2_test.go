package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDirectory(t *testing.T) {
	assert.Equal(t, []int{0, 2, 6}, deleteDirectory(
		[][2]int{{8, 6}, {10, 8}, {6, 0}, {20, 8}, {2, 6}}, 8,
	))
}

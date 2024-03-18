package permute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonRepeatNonOrder(t *testing.T) {
	{
		res := NonRepeatNonOrder([]int{1, 2, 3})
		assert.Equal(t, 6, len(res))
	}

	{
		res := NonRepeatNonOrder([]int{1, 2, 3, 4})
		assert.Equal(t, 24, len(res))
	}
}

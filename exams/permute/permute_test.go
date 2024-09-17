package permute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	{
		res := permute([]int{1, 2, 3})
		assert.Len(t, res, 6)
	}

	{
		res := permute([]int{1, 2, 3, 4})
		assert.Len(t, res, 24)
	}
}

func TestPermuteBacktrace(t *testing.T) {
	{
		res := permuteBacktrace([]int{1, 2, 3})
		assert.Len(t, res, 6)
	}

	{
		res := permuteBacktrace([]int{1, 2, 3, 4})
		assert.Len(t, res, 24)
	}
}

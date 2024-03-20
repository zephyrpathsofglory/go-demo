package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubString(t *testing.T) {
	{
		assert.Equal(t, 5, findSubString("aabeebuu", 1))
		assert.Equal(t, 0, findSubString("baddefguu", 1))
		assert.Equal(t, 0, findSubString("aaaabb", 1))
		assert.Equal(t, 12, findSubString("bbacfgeeiipppeeeee", 3))
		assert.Equal(t, 12, findSubString("bbacfgeeiipppeeeeewwwe", 3))
		assert.Equal(t, 4, findSubString("bbbbbbaaaa", 0))
		assert.Equal(t, 4, findSubString("aaaabbbbb", 0))
		assert.Equal(t, 0, findSubString("aaaabuu", 3))
	}
}

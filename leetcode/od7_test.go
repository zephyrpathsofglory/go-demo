package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinString(t *testing.T) {
	{
		assert.Equal(t, "acdefb", minString("bcdefa"))
		assert.Equal(t, "bbbbbb", minString("bbbbbb"))
		assert.Equal(t, "aaaaabbb", minString("aaaaabbb"))
		assert.Equal(t, "aaaaaab", minString("aaaabaa"))
		assert.Equal(t, "abedcf", minString("afedcb"))
		assert.Equal(t, "abcdef", minString("abcdef"))
		assert.Equal(t, "abzsfw", minString("zbasfw"))
		assert.Equal(t, "afdcafce", minString("efdcafca"))
		assert.Equal(t, "aabbdef", minString("aabbdfe"))
		assert.Equal(t, "aabbeizk", minString("aabbkize"))
	}
}

func TestMinString2(t *testing.T) {
	{
		assert.Equal(t, "acdefb", minString2("bcdefa"))
		assert.Equal(t, "bbbbbb", minString2("bbbbbb"))
		assert.Equal(t, "aaaaabbb", minString2("aaaaabbb"))
		assert.Equal(t, "aaaaaab", minString2("aaaabaa"))
		assert.Equal(t, "abedcf", minString2("afedcb"))
		assert.Equal(t, "abcdef", minString2("abcdef"))
		assert.Equal(t, "abzsfw", minString2("zbasfw"))
		assert.Equal(t, "afdcafce", minString2("efdcafca"))
		assert.Equal(t, "aabbdef", minString2("aabbdfe"))
		assert.Equal(t, "aabbeizk", minString2("aabbkize"))
	}
}

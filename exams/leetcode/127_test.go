package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLadderLength(t *testing.T) {
	assert.Equal(t, 5, ladderLength("hit", "cog", []string{
		"hot", "dot", "dog", "lot", "log", "cog",
	}))
}

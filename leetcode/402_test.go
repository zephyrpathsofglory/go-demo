package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveKdigits(t *testing.T) {
	assert.Equal(t, "157", removeKdigits("76484157", 5))
	assert.Equal(t, "2", removeKdigits("1000234", 3))
	assert.Equal(t, "23", removeKdigits("1000234", 2))
	assert.Equal(t, "0", removeKdigits("00012", 2))
	assert.Equal(t, "401", removeKdigits("45601", 2))
}

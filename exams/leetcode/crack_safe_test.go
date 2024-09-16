package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrackSafe(t *testing.T) {
	assert.Equal(t, "01100", crackSafe(2, 2))
	assert.Equal(t, "0011101000", crackSafe(3, 2))
}

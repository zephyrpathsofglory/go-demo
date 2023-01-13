package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrackSafe(t *testing.T) {
	assert.Equal(t, crackSafe(2, 2), "01100")
	assert.Equal(t, crackSafe(3, 2), "01100")
}

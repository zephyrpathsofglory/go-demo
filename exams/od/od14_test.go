package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSubstring(t *testing.T) {
	assert.Equal(t, 4, isValidSubString("ace", "abcde"))
	assert.Equal(t, -1, isValidSubString("ace", "fsgasd"))
	assert.Equal(t, 2, isValidSubString("ace", "ace"))
	assert.Equal(t, 3, isValidSubString("ace", "ac2e"))
	assert.Equal(t, 3, isValidSubString("ace", "ac2eeee"))
	assert.Equal(t, 0, isValidSubString("a", "ac2eeee"))
	assert.Equal(t, 1, isValidSubString("a", "1a"))
	assert.Equal(t, 1, isValidSubString("a", "1aaa"))
	assert.Equal(t, 6, isValidSubString("aaa", "a1bfdaaaa"))
	assert.Equal(t, 5, isValidSubString("acec", "abcdec"))
}

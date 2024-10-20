package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberToStr(t *testing.T) {
	assert.Equal(t, "twenty two", numberToStr(22))
	assert.Equal(t, "one million one thousand one", numberToStr(1001001))
}

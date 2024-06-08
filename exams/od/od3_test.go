package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideMooncake(t *testing.T) {
	assert.Equal(t, 6, divideMooncake(12, 3))
	assert.Equal(t, 2, divideMooncake(4, 2))
	assert.Equal(t, 2, divideMooncake(5, 3))
}

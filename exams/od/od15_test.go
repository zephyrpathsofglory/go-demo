package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPager(t *testing.T) {
	assert.Equal(t, "1 2 3 ... 10", pager(10, 5, 3))
	assert.Equal(t, "1 2 3 4 5 6 ... 10", pager(10, 8, 3))
	assert.Equal(t, "1 ... 6 7 8 9 10 ... 100", pager(100, 9, 8))
	assert.Equal(t, "1 ... 6 7 8 9 ... 100", pager(100, 8, 8))
	assert.Equal(t, "1 2 3 4 5 6 ... 100", pager(100, 8, 2))
	assert.Equal(t, "1 ... 5 6 7 8 9 10", pager(10, 8, 7))
	assert.Equal(t, "1 ... 95 96 97 98 99 100", pager(100, 8, 98))
}

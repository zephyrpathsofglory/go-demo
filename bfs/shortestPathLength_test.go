package bfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestM(t *testing.T) {
	i := shortestPathLength([][]int{{1, 2, 3}, {0}, {0}, {0}})
	assert.Equal(t, 4, i)
}

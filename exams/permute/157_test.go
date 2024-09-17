package permute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoodsOrder(t *testing.T) {
	assert.Len(t, goodsOrder("wage"), 24)
}

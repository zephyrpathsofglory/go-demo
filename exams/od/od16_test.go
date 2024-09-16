package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrEdit(t *testing.T) {
	res := strEdit("aaaa+bbbb~@cc<<<^--d@d")
	assert.Equal(t, []string{"aaDd", "bbbCC"}, res)
}

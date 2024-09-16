package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitString(t *testing.T) {
	assert.Equal(t, "12abc-abc-ABC-4aB-@", SplitString("12abc-abCABc-4aB@", 3))
	assert.Equal(t, "12abc-abCABc4aB@", SplitString("12abc-abCABc-4aB@", 12))
}

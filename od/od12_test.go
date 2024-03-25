package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecompression(t *testing.T) {
	assert.Equal(t, "kkkmnmn", decompression("3[k]2[mn]"))
	assert.Equal(t, "mccmccmcc", decompression("3[m2[c]]"))
	assert.Equal(t, "hmccmccmcc", decompression("h3[m2[c]]"))
	assert.Equal(t, "hmccmccmccg", decompression("h3[m2[c]]g"))
}

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
	assert.Equal(t, "kkmnmn", decompression("2[k]2[mn]"))
	assert.Equal(t, "mccmcc", decompression("2[m2[c]]"))
	assert.Equal(t, "wc", decompression("1[w]c"))
	assert.Equal(t, "wcwc", decompression("2[wc]"))
	assert.Equal(t, "wkkkcwkkkc", decompression("2[w3[k]c]"))
}

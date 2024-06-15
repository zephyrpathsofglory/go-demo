package apitest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLParse(t *testing.T) {
	{
		h, p, err := GetHostAndPortFromURL("schme://user:passwd@127.0.0.1:1777?sslmode=false")
		assert.Nil(t, err)
		assert.Equal(t, "127.0.0.1", h)
		assert.Equal(t, "1777", p)
	}
	// {
	// 	h, p, err := GetHostAndPortFromURL("schme://user/passwd@127.0.0.1:1777?sslmode=false")
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, "127.0.0.1", h)
	// 	assert.Equal(t, "1777", p)
	// }
	// {
	// 	h, p, err := GetHostAndPortFromURL("user/passwd@127.0.0.1:1777?sslmode=false")
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, "127.0.0.1", h)
	// 	assert.Equal(t, "1777", p)
	// }
}

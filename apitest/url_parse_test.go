package apitest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLParse(t *testing.T) {
	h, p, err := GetHostAndPortFromURL("schme://user:passwd@127.0.0.1:1777?sslmode=false")
	assert.Nil(t, err)
	assert.Equal(t, h, "127.0.0.1")
	assert.Equal(t, p, "1777")
}

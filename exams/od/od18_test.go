package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSameSubNet(t *testing.T) {
	assert.Equal(t, uint8(2), inSameSubNet("192.168.0.2", "192.168.0.253", "255.255.255.128"))
	assert.Equal(t, uint8(0), inSameSubNet("192.168.0.1", "192.168.0.254", "255.255.255.0"))
	assert.Equal(t, uint8(1), inSameSubNet("192.168.0.1", "192.168.0.254", "255.225.255.0"))
	assert.Equal(t, uint8(1), inSameSubNet("192.168.10.4", "192.168.224.256", "255.255.255.0"))
	assert.Equal(t, uint8(2), inSameSubNet("232.43.7.59", "193.194.202.15", "255.0.0.0"))
}

package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	assert.Equal(t, 3/2, divide(3, 2))
	assert.Equal(t, 6/3, divide(6, 3))
	assert.Equal(t, 2147483647, divide(6, 0))
	assert.Equal(t, -6/3, divide(-6, 3))
	assert.Equal(t, 6/-3, divide(6, -3))
	assert.Equal(t, -2147483647, divide(2147483647, -1))

	// int32 overflow
	assert.Equal(t, 2147483647, divide(-2147483648, -1))
}

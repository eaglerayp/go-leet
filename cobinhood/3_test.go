package cobinhood

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFab(t *testing.T) {
	assert.Equal(t, 0, Fibonacci(0))
	assert.Equal(t, 1, Fibonacci(1))
	assert.Equal(t, 1, Fibonacci(2))
	assert.Equal(t, 2, Fibonacci(3))
	assert.Equal(t, 3, Fibonacci(4))
	assert.Equal(t, 5, Fibonacci(5))
	assert.Equal(t, 8, Fibonacci(6))
	assert.Equal(t, 13, Fibonacci(7))
	assert.Equal(t, 21, Fibonacci(8))
	assert.Equal(t, 34, Fibonacci(9))
	assert.Equal(t, 227465, Fibonacci(35))
	assert.Equal(t, 930352, Fibonacci(36))
	assert.Equal(t, 915075, Fibonacci(100))
}

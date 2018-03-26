package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAndSay(t *testing.T) {
	assert.Equal(t, "1", countAndSay(1))
	assert.Equal(t, "11", countAndSay(2))
	assert.Equal(t, "21", countAndSay(3))
	assert.Equal(t, "1211", countAndSay(4))
	assert.Equal(t, "111221", countAndSay(5))
	assert.Equal(t, "312211", countAndSay(6))
	assert.Equal(t, "13112221", countAndSay(7))
}

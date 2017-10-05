package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	assert.Equal(t, -1, strStr("123", "456"))
	assert.Equal(t, 3, strStr("123456", "456"))
	assert.Equal(t, 0, strStr("456", "456"))
	assert.Equal(t, -1, strStr("123", "4567"))
	assert.Equal(t, 0, strStr("456", ""))
	assert.Equal(t, 0, strStr("", ""))
}

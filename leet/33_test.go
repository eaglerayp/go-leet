package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRotate(t *testing.T) {
	assert.Equal(t, 3, search([]int{9, 12, 13, 1, 6, 8}, 1))
	assert.Equal(t, -1, search([]int{9, 12, 13, 1, 6, 8}, 2))
	assert.Equal(t, 1, search([]int{9, 12, 13, 1, 6, 8}, 12))
	assert.Equal(t, 3, search([]int{9, 12, 13, 6, 8}, 6))
	assert.Equal(t, -1, search([]int{}, 2))
	assert.Equal(t, 0, search([]int{2}, 2))
	assert.Equal(t, 0, search([]int{2, 1}, 2))
	assert.Equal(t, 1, search([]int{2, 1}, 1))
}

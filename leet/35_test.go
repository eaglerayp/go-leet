package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	assert.Equal(t, 0, searchInsert([]int{1, 6, 8, 9, 13}, 1))
	assert.Equal(t, 4, searchInsert([]int{1, 6, 8, 9, 13}, 13))
	assert.Equal(t, 5, searchInsert([]int{1, 6, 8, 9, 10, 13}, 13))
	assert.Equal(t, 2, searchInsert([]int{1, 6, 8, 9, 10, 13}, 7))
	assert.Equal(t, 0, searchInsert([]int{}, 2))
	assert.Equal(t, 0, searchInsert([]int{2}, 2))
	assert.Equal(t, 0, searchInsert([]int{2}, 1))
	assert.Equal(t, 1, searchInsert([]int{2}, 3))
	assert.Equal(t, 0, searchInsert([]int{0, 1}, 0))
	assert.Equal(t, 1, searchInsert([]int{0, 1}, 1))
	assert.Equal(t, 2, searchInsert([]int{0, 1}, 51))
	assert.Equal(t, 1, searchInsert([]int{1, 4}, 4))
}

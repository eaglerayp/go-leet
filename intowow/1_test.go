package intowow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindListLength(t *testing.T) {
	assert.Equal(t, 4, FindListLength([]int{1, 4, -1, 3, 2}))
	assert.Equal(t, 1, FindListLength([]int{-1}))
	assert.Equal(t, 2, FindListLength([]int{1, -1}))
	assert.Equal(t, 12, FindListLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, -1, 12, 13, 14, 15}))
	assert.Equal(t, 16, FindListLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, -1}))
}

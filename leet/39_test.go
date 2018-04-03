package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	result := combinationSum([]int{2, 3, 6, 7}, 7)
	assert.Equal(t, []int{2, 2, 3}, result[0])
	assert.Equal(t, []int{7}, result[1])
	result = combinationSum([]int{2, 3, 6, 7}, 9)
	assert.Equal(t, []int{2, 2, 2, 3}, result[0])
	assert.Equal(t, []int{2, 7}, result[1])
	assert.Equal(t, []int{3, 3, 3}, result[2])
	assert.Equal(t, []int{3, 6}, result[3])
	result = combinationSum([]int{2, 3, 6, 7}, 1)
	assert.Nil(t, result)
}

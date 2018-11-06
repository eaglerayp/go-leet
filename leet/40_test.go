package leet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum2(t *testing.T) {
	result := combinationSum2([]int{2, 3, 6, 7}, 7)
	assert.Equal(t, []int{7}, result[0])
	result = combinationSum2([]int{2, 2, 3, 6, 7}, 7)
	assert.Equal(t, []int{2, 2, 3}, result[0])
	assert.Equal(t, []int{7}, result[1])
	result = combinationSum2([]int{2, 3, 6, 7}, 9)
	assert.Equal(t, []int{2, 7}, result[0])
	assert.Equal(t, []int{3, 6}, result[1])
	result = combinationSum2([]int{2, 3, 6, 7}, 1)
	assert.Nil(t, result)
	result = combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Println(result)
	assert.Equal(t, []int{1, 7}, result[2])
	assert.Equal(t, []int{1, 2, 5}, result[1])
	assert.Equal(t, []int{2, 6}, result[3])
	assert.Equal(t, []int{1, 1, 6}, result[0])
}

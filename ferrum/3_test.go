package ferrum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMinimumUniqueSum(t *testing.T) {
	// Write your code here
	assert.Equal(t, int32(3), getMinimumUniqueSum([]int32{1, 1}))
	assert.Equal(t, int32(6), getMinimumUniqueSum([]int32{1, 2, 2}))
	assert.Equal(t, int32(9), getMinimumUniqueSum([]int32{2, 2, 2}))
	assert.Equal(t, int32(17), getMinimumUniqueSum([]int32{3, 2, 1, 2, 7}))
}

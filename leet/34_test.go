package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRange(t *testing.T) {
	t1 := searchRange([]int{1, 6, 6, 8, 9}, 6)
	assert.Equal(t, 1, t1[0])
	assert.Equal(t, 2, t1[1])
	t2 := searchRange([]int{1, 6, 6, 8, 9}, 2)
	assert.Equal(t, -1, t2[0])
	assert.Equal(t, -1, t2[1])
	t3 := searchRange([]int{}, 2)
	assert.Equal(t, -1, t3[0])
	assert.Equal(t, -1, t3[1])
	t4 := searchRange([]int{2}, 2)
	assert.Equal(t, 0, t4[0])
	assert.Equal(t, 0, t4[1])
	t5 := searchRange([]int{1, 5}, 1)
	assert.Equal(t, 0, t5[0])
	assert.Equal(t, 0, t5[1])
	t6 := searchRange([]int{1, 2, 4, 6, 6, 6, 6, 6, 6, 6, 8, 9}, 6)
	assert.Equal(t, 3, t6[0])
	assert.Equal(t, 9, t6[1])
	t7 := searchRange([]int{5, 5}, 5)
	assert.Equal(t, 0, t7[0])
	assert.Equal(t, 1, t7[1])
	t8 := searchRange([]int{5, 5, 5, 5, 5, 5, 5, 5}, 5)
	assert.Equal(t, 0, t8[0])
	assert.Equal(t, 7, t8[1])
	t9 := searchRange([]int{1, 5}, 5)
	assert.Equal(t, 1, t9[0])
	assert.Equal(t, 1, t9[1])
}

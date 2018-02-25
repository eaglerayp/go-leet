package intowow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrossRiver(t *testing.T) {
	assert.Equal(t, 2, CrossRiver([]int{1, -1, 0, 2, 3, 5}, 3))

	assert.Equal(t, 0, CrossRiver([]int{-1, -1}, 3))
	assert.Equal(t, 1, CrossRiver([]int{-1, 1, -1}, 2))
	assert.Equal(t, -1, CrossRiver([]int{-1, -1}, 1))
	assert.Equal(t, -1, CrossRiver([]int{1, -1, 0, -1, -1, -1}, 3))
}

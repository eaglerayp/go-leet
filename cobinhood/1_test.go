package cobinhood

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCastle(t *testing.T) {
	a1 := Solution([]int{-3, -3})
	assert.Equal(t, 1, a1)
	a2 := Solution([]int{2, 2, 3, 4, 3, 3, 2, 2, 1, 1, 2, 5})
	assert.Equal(t, 4, a2)
	a3 := Solution([]int{-3, 3})
	assert.Equal(t, 2, a3)
	a4 := Solution([]int{})
	assert.Equal(t, 0, a4)
}

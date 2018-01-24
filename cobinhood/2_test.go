package cobinhood

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdjacnetMinDistance(t *testing.T) {
	// a1 := AdjacentMinimumDistance([]int{0, 3, 3, 7, 5, 3, 11, 1})
	// assert.Equal(t, 0, a1)
	// a2 := AdjacentMinimumDistance([]int{0, 3, 7, 5, 11, 1})
	// assert.Equal(t, 2, a2)
	// a3 := AdjacentMinimumDistance([]int{-3, 3})
	// assert.Equal(t, 6, a3)
	// a4 := AdjacentMinimumDistance([]int{})
	// assert.Equal(t, -1, a4)
	a5 := AdjacentMinimumDistance([]int{MIN_INT, MAX_INT})
	assert.Equal(t, -2, a5)
	// fmt.Println("Dif:", MAX_INT-MIN_INT)
}

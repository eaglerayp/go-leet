package intowow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuelQueue(t *testing.T) {
	assert.Equal(t, 8, FuelQueue([]int{2, 8, 4, 3, 2}, 7, 11, 3))
	assert.Equal(t, -1, FuelQueue([]int{5}, 4, 0, 3))
	assert.Equal(t, -1, FuelQueue([]int{1, 5}, 4, 0, 3))
}

package intowow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCountry(t *testing.T) {
	i1 := [][]int{}
	i1 = append(i1, []int{5, 4, 4})
	i1 = append(i1, []int{4, 3, 4})
	i1 = append(i1, []int{3, 2, 4})
	i1 = append(i1, []int{2, 2, 2})
	i1 = append(i1, []int{3, 3, 4})
	i1 = append(i1, []int{1, 4, 4})
	i1 = append(i1, []int{4, 1, 1})
	assert.Equal(t, 11, CountCountry(i1))
	i1[6][2] = 4
	assert.Equal(t, 11, CountCountry(i1))
	i1[6][2] = 2
	assert.Equal(t, 12, CountCountry(i1))
	i1[6] = []int{4, 4, 4}
	assert.Equal(t, 9, CountCountry(i1))
}

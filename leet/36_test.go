package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSudoku(t *testing.T) {
	i1 := [][]byte{}
	i1 = append(i1, []byte("825471396"))
	i1 = append(i1, []byte("194326578"))
	i1 = append(i1, []byte("376985241"))
	i1 = append(i1, []byte("519743862"))
	i1 = append(i1, []byte("632598417"))
	i1 = append(i1, []byte("487612935"))
	i1 = append(i1, []byte("263159784"))
	i1 = append(i1, []byte("948267153"))
	i1 = append(i1, []byte("751834629"))

	assert.True(t, isValidSudoku(i1))

	i1 = [][]byte{}
	i1 = append(i1, []byte("825471396"))
	i1 = append(i1, []byte("194326578"))
	i1 = append(i1, []byte("376985241"))
	i1 = append(i1, []byte("519743862"))
	i1 = append(i1, []byte("632598417"))
	i1 = append(i1, []byte("487612935"))
	i1 = append(i1, []byte("263159784"))
	i1 = append(i1, []byte("948267153"))
	i1 = append(i1, []byte("751834625"))
	assert.False(t, isValidSudoku(i1))

	i1 = [][]byte{}
	i1 = append(i1, []byte("825471396"))
	i1 = append(i1, []byte("194326578"))
	i1 = append(i1, []byte("376985241"))
	i1 = append(i1, []byte("519743862"))
	i1 = append(i1, []byte("632598417"))
	i1 = append(i1, []byte("487612935"))
	i1 = append(i1, []byte("263159784"))
	i1 = append(i1, []byte("948267153"))
	i1 = append(i1, []byte("75183462."))
	assert.True(t, isValidSudoku(i1))

	i1 = [][]byte{}
	i1 = append(i1, []byte(".25471396"))
	i1 = append(i1, []byte("194326578"))
	i1 = append(i1, []byte("376985241"))
	i1 = append(i1, []byte("5197.3862"))
	i1 = append(i1, []byte("632598417"))
	i1 = append(i1, []byte("487612935"))
	i1 = append(i1, []byte("2..159784"))
	i1 = append(i1, []byte("94826...3"))
	i1 = append(i1, []byte("75183462."))
	assert.True(t, isValidSudoku(i1))
}

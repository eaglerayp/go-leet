package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	assert.Equal(t, "()", generateParenthesis(1)[0])
	n2 := generateParenthesis(2)
	assert.Equal(t, "()()", n2[0])
	assert.Equal(t, "(())", n2[1])
}

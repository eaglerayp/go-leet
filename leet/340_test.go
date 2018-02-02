package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	assert.Equal(t, 1, lengthOfLongestSubstringKDistinct("1", 1))
	assert.Equal(t, 6, lengthOfLongestSubstringKDistinct("abcbcad", 3))
	assert.Equal(t, 4, lengthOfLongestSubstringKDistinct("abacacd", 2))
	assert.Equal(t, 8, lengthOfLongestSubstringKDistinct("abcbcadaccaa", 3))
}

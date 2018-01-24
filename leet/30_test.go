package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubstring(t *testing.T) {
	r1 := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	assert.Equal(t, 2, len(r1))
	assert.Equal(t, 0, r1[0])
	assert.Equal(t, 9, r1[1])

	r2 := findSubstring("bar", []string{"foo", "bar"})
	assert.Equal(t, 0, len(r2))

	r3 := findSubstring("foofoobarfoofoo", []string{"foo", "bar", "foo"})
	assert.Equal(t, 3, len(r3))
	assert.Equal(t, 0, r3[0])
	assert.Equal(t, 3, r3[1])
	assert.Equal(t, 6, r3[2])

	r4 := findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})
	assert.Equal(t, 3, len(r4))
	assert.Equal(t, 6, r4[0])
	assert.Equal(t, 9, r4[1])
	assert.Equal(t, 12, r4[2])
}

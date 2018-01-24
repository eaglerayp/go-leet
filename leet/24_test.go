package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	e342 := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	n := swapPairs(e342)
	assert.Equal(t, 3, n.Val)
	n = n.Next
	assert.Equal(t, 2, n.Val)
	n = n.Next
	assert.Equal(t, 4, n.Val)

	n = swapPairs(nil)
	assert.Nil(t, n)

	n = swapPairs(&ListNode{Val: 2})
	assert.Equal(t, 2, n.Val)

	n = swapPairs(&ListNode{Val: 2, Next: &ListNode{Val: 3}})
	assert.Equal(t, 3, n.Val)
	n = n.Next
	assert.Equal(t, 2, n.Val)
}

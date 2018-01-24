package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLenLarger(t *testing.T) {
	e234 := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	assert.True(t, IsLenLarger(e234, 2))
	assert.False(t, IsLenLarger(e234, 3))
}

func TestReverseKGroup(t *testing.T) {
	e234 := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	cur := reverseKGroup(e234, 2)
	assert.Equal(t, 3, cur.Val)
	cur = cur.Next
	assert.Equal(t, 2, cur.Val)
	cur = cur.Next
	assert.Equal(t, 4, cur.Val)

	e234 = &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	cur = reverseKGroup(e234, 3)
	assert.Equal(t, 4, cur.Val)
	cur = cur.Next
	assert.Equal(t, 3, cur.Val)
	cur = cur.Next
	assert.Equal(t, 2, cur.Val)

	e234 = &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	cur = reverseKGroup(e234, 4)
	assert.Equal(t, 2, cur.Val)
	cur = cur.Next
	assert.Equal(t, 3, cur.Val)
	cur = cur.Next
	assert.Equal(t, 4, cur.Val)

	e23 := &ListNode{Val: 2, Next: &ListNode{Val: 3}}
	cur = reverseKGroup(e23, 2)
	assert.Equal(t, 3, cur.Val)
	cur = cur.Next
	assert.Equal(t, 2, cur.Val)

	e23 = &ListNode{Val: 2, Next: &ListNode{Val: 3}}
	cur = reverseKGroup(e23, 1)
	assert.Equal(t, 2, cur.Val)
	cur = cur.Next
	assert.Equal(t, 3, cur.Val)
}

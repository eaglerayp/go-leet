package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	e342 := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	e765 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}}
	e189 := &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9}}}
	k := []*ListNode{e342, e765, e189}
	n := mergeKLists(k)
	assert.Equal(t, 1, n.Val)
	n = n.Next
	assert.Equal(t, 2, n.Val)
	n = n.Next
	assert.Equal(t, 3, n.Val)
	n = n.Next
	assert.Equal(t, 4, n.Val)
	n = n.Next
	assert.Equal(t, 5, n.Val)
	n = n.Next
	assert.Equal(t, 6, n.Val)
	n = n.Next
	assert.Equal(t, 7, n.Val)
	n = n.Next
	assert.Equal(t, 8, n.Val)
	n = n.Next
	assert.Equal(t, 9, n.Val)

	n = mergeKLists([]*ListNode{e342})
	assert.Equal(t, 2, n.Val)
	n = n.Next
	assert.Equal(t, 3, n.Val)
	n = n.Next
	assert.Equal(t, 4, n.Val)

	n = mergeKLists([]*ListNode{})
	assert.Nil(t, n)

	e := &ListNode{}
	n = mergeKLists([]*ListNode{e})
	assert.Equal(t, 0, n.Val)
	n = n.Next
	assert.Nil(t, n)

	n = mergeKLists([]*ListNode{nil, nil})
	assert.Nil(t, n)
}

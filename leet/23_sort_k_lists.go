package leet

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// An IntHeap is a min-heap of ints.
type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	// using min Heap
	n := len(lists)
	h := &NodeHeap{}
	for i := 0; i < n; i++ {
		if lists[i] != nil {
			*h = append(*h, lists[i])
		}
	}
	if h.Len() < 1 {
		return nil
	}
	heap.Init(h)
	// pop out all heap
	head := heap.Pop(h).(*ListNode)
	if head == nil {
		return head
	}
	cur := head
	if cur.Next != nil {
		heap.Push(h, cur.Next)
	}
	for h.Len() > 0 {
		cur.Next = heap.Pop(h).(*ListNode)
		cur = cur.Next
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
	}
	return head
}

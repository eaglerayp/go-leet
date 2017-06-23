package algo

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	list1 := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}}
	iterateList(list1)
	fmt.Println("t1")
	iterateList(removeNthFromEnd(list1, 1))
	list1 = &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}}
	fmt.Println("t2")
	iterateList(removeNthFromEnd(list1, 2))
	list1 = &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}}
	fmt.Println("t3")
	iterateList(removeNthFromEnd(list1, 3))
}

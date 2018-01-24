package leet

import "fmt"

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// IterList is a helper function for debug
func IterList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

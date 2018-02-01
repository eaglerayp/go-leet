package main

import (
	"fmt"
)

// linked list

// reverse -n node

type Node struct {
	Next *Node
	Val  int
}

func RemoveReverseN(head *Node, N int) error {
	fast := head

	// n = 3, N =1
	for i := 0; i < N; i++ {
		// n =3, N=4
		if fast == nil {
			return fmt.Errorf("invalid N")
		}
		fast = fast.Next
	}

	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return nil
}

func GetListLen(head *Node) int {
	n := 0
	if head == nil {
		return n
	}
	for head.Next != nil {
		head = head.Next
		n++
	}
	return n
}

func main() {

}

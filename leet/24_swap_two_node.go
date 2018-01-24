package leet

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := head.Next
	next := swapPairs(temp.Next)
	temp.Next = head
	head.Next = next
	head = temp
	return head
}

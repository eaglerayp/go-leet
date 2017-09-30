package leet

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	if !IsLenLarger(head, k-1) {
		return head
	}
	prev := head
	cur := prev.Next
	// reverse list
	for i := 0; i < k-1; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	head.Next = reverseKGroup(cur, k)
	return prev
}

func IsLenLarger(head *ListNode, n int) bool {
	if n <= 0 {
		return true
	}
	if head == nil {
		return false
	}
	if head.Next != nil {
		return IsLenLarger(head.Next, n-1)
	}
	return false
}

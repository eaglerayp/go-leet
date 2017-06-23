package algo

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// use two pointer for temp remember Nth before current
	var end, cur, prev *ListNode
	end = head
	// walk n
	for i := 0; i < n; i++ {
		end = end.Next
		// if end == nil during, walk n, panic
	}
	if end == nil {
		return head.Next
	}
	cur = head
	prev = head
	for end != nil {
		end = end.Next
		prev = cur
		cur = cur.Next
	}
	prev.Next = cur.Next
	return head
}

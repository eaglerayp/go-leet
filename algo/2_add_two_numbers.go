package algo

// You are given **two non-empty** linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// first version i choose to change it to int then addTwoNumbers
	// but this question should add digit by digit
	null1 := l1 == nil
	null2 := l2 == nil
	carry := 10
	// this is head (dummy node)
	result := &ListNode{}
	prev := result
	ac := 0
	// from digit one to last
	for !null1 && !null2 {
		digit := l1.Val + l2.Val + ac
		if digit >= carry {
			ac = 1
			digit -= carry
		} else {
			ac = 0
		}
		now := &ListNode{Val: digit}
		prev.Next = now
		l1 = l1.Next
		l2 = l2.Next
		prev = now
		null1 = l1 == nil
		null2 = l2 == nil
	}
	if null1 {
		// traverse l2
		for !null2 {
			digit := l2.Val + ac
			if digit >= carry {
				ac = 1
				digit -= carry
			} else {
				ac = 0
			}
			l2.Val = digit
			prev.Next = l2
			prev = l2
			l2 = l2.Next

			null2 = l2 == nil
		}
		// final add carry
		if ac > 0 {
			now := &ListNode{Val: 1}
			prev.Next = now
		}
	} else if null2 {
		// traverse l1
		for !null1 {
			digit := l1.Val + ac
			if digit >= carry {
				ac = 1
				digit -= carry
			} else {
				ac = 0
			}
			l1.Val = digit
			prev.Next = l1
			prev = l1
			l1 = l1.Next
			null1 = l1 == nil
		}
		// final add carry
		if ac > 0 {
			now := &ListNode{Val: 1}
			prev.Next = now
		}
	}
	return result.Next
}

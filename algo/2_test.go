package algo

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	e342 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	e465 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	e999 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}
	e0 := &ListNode{Val: 0}
	e1 := &ListNode{Val: 1}
	a1 := listToNumber(addTwoNumbers(e342, e465))
	a2 := listToNumber(addTwoNumbers(e0, e0))
	a3 := listToNumber(addTwoNumbers(e342, e0))
	a4 := listToNumber(addTwoNumbers(e999, e1))
	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a3)
	// fmt.Println(a4)
	if a1 != 807 {
		t.Fail()
	}
	if a2 != 0 {
		t.Fail()
	}
	if a3 != 342 {
		t.Fail()
	}
	if a4 != 1000 {
		t.Fail()
	}
}

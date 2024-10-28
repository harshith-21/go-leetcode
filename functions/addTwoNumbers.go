package functions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	current := ans
	carry := 0
	for l1 != nil || l2 != nil {
		l1val, l2val := 0, 0
		if l1 != nil {
			l1val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2val = l2.Val
			l2 = l2.Next
		}
		sum := (l1val + l2val + carry) % 10
		carry = (l1val + l2val + carry) / 10
		newNode := &ListNode{Val: sum}
		current.Next = newNode
		current = current.Next
	}
	if carry != 0 {
		newNode := &ListNode{Val: carry}
		current.Next = newNode
		current = current.Next
	}
	return ans.Next
}

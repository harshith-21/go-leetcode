package functions

func middleNode(head *ListNode) *ListNode {
	tempslow := head
	tempfast := head

	for tempfast != nil && tempfast.Next != nil {
		tempfast = tempfast.Next.Next
		tempslow = tempslow.Next
	}
	return tempslow
}

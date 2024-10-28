package functions

func hasCycle(head *ListNode) bool {
	tempslow := head
	tempfast := head

	for tempfast != nil && tempfast.Next != nil {
		tempslow = tempslow.Next
		tempfast = tempfast.Next.Next
		if tempfast == tempslow {
			return true
		}
	}
	return false
}

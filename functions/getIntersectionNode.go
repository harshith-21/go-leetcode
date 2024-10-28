package functions

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodelist := map[*ListNode]int{}

	tempa := headA
	tempb := headB

	for tempa != nil || tempb != nil {
		if tempa != nil {
			_, oka := nodelist[tempa]
			if oka {
				return tempa
			} else {
				nodelist[tempa] = 1
			}
			tempa = tempa.Next
		}
		if tempb != nil {
			_, okb := nodelist[tempb]
			if okb {
				return tempb
			} else {
				nodelist[tempb] = 1
			}
			tempb = tempb.Next
		}
	}
	// this wont happen, will be returned from above itself
	return tempa
}

package functions

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ans := &ListNode{}
	temp := ans

	//l1 := list1
	//l2 := list2

	for list1 != nil && list2 != nil {

		if list1.Val > list2.Val {
			temp.Next = list1
			list1 = list1.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
		}
		temp = temp.Next
	}
	if list1 == nil {
		temp.Next = list2
	} else {
		temp.Next = list1
	}
	return ans.Next
}

//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	ans := &ListNode{}
//	temp := ans
//
//	l1 := list1
//	l2 := list2
//
//	for l1 != nil && l2 != nil {
//		newnode := &ListNode{}
//		if l1.Val > l2.Val {
//			newnode.Val = l2.Val
//			l2 = l2.Next
//		} else {
//			newnode.Val = l1.Val
//			l1 = l1.Next
//		}
//		temp.Next = newnode
//		temp = temp.Next
//	}
//	if l1 == nil {
//		temp.Next = l2
//	} else {
//		temp.Next = l1
//	}
//	return ans.Next
//}

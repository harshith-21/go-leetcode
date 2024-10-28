package functions

import "fmt"

func ModifiedList(nums []int, head *ListNode) *ListNode {
	headtemp := head

	//hashmap for faster lookup instead of contains
	numMap := make(map[int]struct{})
	for _, num := range nums {
		numMap[num] = struct{}{}
	}

	for headtemp != nil {
		for contains(numMap, headtemp.Val) {
			head = head.Next
			headtemp = head
		}
		for headtemp != nil && headtemp.Next != nil {
			if contains(numMap, headtemp.Next.Val) {
				headtemp.Next = headtemp.Next.Next
			} else {
				headtemp = headtemp.Next
			}
		}
		if headtemp != nil {
			headtemp = headtemp.Next
		}
	}
	return head
}

func contains(numMap map[int]struct{}, n int) bool {
	//for _, num := range nums {
	//	if num == n {
	//		return true
	//	}
	//}
	//return false
	_, exists := numMap[n]
	return exists
}

func TestModifiedList(nums []int, heads []int) {
	head := &ListNode{Val: heads[0]}
	current := head
	for _, val := range heads[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	currenttep := ModifiedList(nums, head)
	for currenttep != nil {
		fmt.Print(currenttep.Val, " ")
		currenttep = currenttep.Next
	}
	fmt.Println()
}

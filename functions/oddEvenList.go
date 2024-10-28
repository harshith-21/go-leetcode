package functions

import "fmt"

func oddEvenList(head *ListNode) *ListNode {
	if !isLengthgt1(head) {
		return head
	}
	chainodd := head
	chaineven := head.Next
	evenHead := chaineven
	oddhead := chainodd

	for chaineven != nil && chaineven.Next != nil {
		chainodd.Next = chaineven.Next
		chainodd = chainodd.Next
		chaineven.Next = chainodd.Next
		chaineven = chaineven.Next
	}
	PrintListlimit(evenHead, 100)
	PrintListlimit(oddhead, 100)

	chainodd.Next = evenHead

	return oddhead
}

func isLengthgt1(head *ListNode) bool {
	i := 0
	for head != nil {
		head = head.Next
		i++
		if i > 1 {
			return true
		}
	}
	return false
}

func TestoddEvenList() {
	list := []int{1, 2, 3, 4, 5}
	head := CreateList(list)
	PrintListlimit(head, 5)
	newHead := oddEvenList(head)
	PrintListlimit(newHead, 5)
}

func PrintListlimit(head *ListNode, count int) {
	current := head
	for current != nil && count > 0 {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
		count--
	}
	fmt.Println("nil")
}

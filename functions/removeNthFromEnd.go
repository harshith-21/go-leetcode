package functions

import "fmt"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headnew := &ListNode{0, head}
	fmt.Println("headnew", headnew)
	temp := headnew
	tempslow := headnew
	for i := 0; i < n; i++ {
		temp = temp.Next
		fmt.Println("loop")
	}
	fmt.Println("temp", temp)
	fmt.Println("tempslow", tempslow)
	for temp != nil && temp.Next != nil {
		temp = temp.Next
		tempslow = tempslow.Next
	}
	fmt.Println("temp", temp)
	fmt.Println("tempslow", tempslow)
	tempslow.Next = tempslow.Next.Next
	return headnew.Next
}

func removeNthFromEnd_backup(head *ListNode, n int) *ListNode {
	temp := head
	tempslow := head

	for n != 0 && temp != nil && temp.Next != nil {
		temp = temp.Next
		n--
	}
	for temp != nil && temp.Next != nil {
		temp = temp.Next
		tempslow = tempslow.Next
	}
	tempslow.Next = tempslow.Next.Next
	return head
}

func CreateList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head

	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

func PrintList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func TestremoveNthFromEnd() {
	//listarr := []int{1, 2}
	//head := CreateList(listarr)
	//newhead := removeNthFromEnd(head, 1)
	//PrintList(newhead)

	listarr2 := []int{1}
	head2 := CreateList(listarr2)
	newhead2 := removeNthFromEnd(head2, 1)
	PrintList(newhead2)
}

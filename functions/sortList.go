package functions

import "fmt"

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length := findLength(head)
	return mergesort(head, length)
}

func mergeList(head1 *ListNode, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			current.Next = head1
			head1 = head1.Next
		} else {
			current.Next = head2
			head2 = head2.Next
		}
		current = current.Next
	}
	if head1 != nil {
		current.Next = head1
	}
	if head2 != nil {
		current.Next = head2
	}
	return dummy.Next
}

func mergesort(head *ListNode, length int) *ListNode {
	if length <= 1 {
		return head
	}

	middle := findMiddle(head)
	left := head
	right := middle.Next
	middle.Next = nil // Split the list

	leftSorted := mergesort(left, length/2)
	rightSorted := mergesort(right, length-length/2)

	return mergeList(leftSorted, rightSorted)
}

func findLength(head *ListNode) int {
	length := 0
	for head != nil {
		head = head.Next
		length++
	}
	return length
}

func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// `prev` will be the node just before the middle of the list
	if prev != nil {
		return prev
	}
	return head
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func TestSortList() {
	list := []int{-1, 5, 3, 4, 0}
	head := CreateList(list)
	PrintList(head)
	newHead := sortList(head)
	PrintList(newHead)
}

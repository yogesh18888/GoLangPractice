package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	list1 := CreateList([]int{1, 2, 4})
	list2 := CreateList([]int{1, 3, 4})

	fmt.Println("List 1:")
	PrintList(list1)

	fmt.Println("List 2:")
	PrintList(list2)

	// Merge and print result
	merged := mergeTwoLists(list1, list2)

	fmt.Println("Merged List:")
	PrintList(merged)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		fmt.Println(fmt.Sprintf("%d : %d", list1.Val, list2.Val))
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head
	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper to print a linked list
func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

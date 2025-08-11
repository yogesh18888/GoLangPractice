package main

import (
	"fmt"
	"sort"
)

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
	merged := mergeTwoLists1(list1, list2)

	fmt.Println("Merged List:")
	PrintList(merged)
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {

	slice1 := listNodeToList(list1)
	fmt.Println("slice1:", slice1)
	slice2 := listNodeToList(list2)
	fmt.Println("slice2:", slice2)

	resultSlice := append(slice1, slice2...)
	fmt.Println("resultSlice:", resultSlice)

	sort.Ints(resultSlice)
	fmt.Println("sorted resultSlice:", resultSlice)

	return listToListNode(resultSlice)
}

func listNodeToList(listNode *ListNode) []int {
	tmpList := listNode
	resultList := []int{}
	for {
		if tmpList == nil {
			fmt.Println("break")
			break
		}
		resultList = append(resultList, tmpList.Val)
		tmpList = tmpList.Next
	}
	return resultList
}

func listToListNode(resultList []int) *ListNode {
	var listNode *ListNode
	lastNode := &ListNode{}
	for _, val := range resultList {
		newListNode := &ListNode{
			Val:  val,
			Next: &ListNode{},
		}
		if listNode == nil {
			lastNode = newListNode
			listNode = lastNode
			continue
		}
		lastNode.Next = newListNode
		lastNode = newListNode
	}
	lastNode.Next = nil
	return listNode
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

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	seenMap := make(map[int]bool)
	for _, value := range nums {
		if seenMap[value] {
			return true
		}
		seenMap[value] = true
	}
	return false
}

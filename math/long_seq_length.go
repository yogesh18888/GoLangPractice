package main

import (
	"fmt"
	"sort"
)

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
// You must write an algorithm that runs in O(n) time.
func main() {
	nums := []int{100, 4, 200, 1, 1, 3, 2, 505, 504, 503, 502, 501, 500}
	sort.Ints(nums)
	fmt.Println("sorted:", nums)
	currentLength := 1
	maxLength := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		} else if nums[i] == nums[i-1]+1 {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 1
		}
	}
	if currentLength > maxLength {
		maxLength = currentLength
	}
	fmt.Println("max length:", maxLength)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	array := [5]int{4, 3, 5, 1, 2}
	fmt.Println("array before sort:", array)
	arraySlice := array[:]
	sort.Ints(arraySlice)
	fmt.Println("sorted arraySlice:", arraySlice)
	fmt.Println("sorted array:", array) //since slice works on array copy it indirectly sorts array too
}

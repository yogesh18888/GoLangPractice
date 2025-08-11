package main

import "fmt"

func main() {
	originalSlice := []string{"A", "B", "C", "D", "E"}
	// Omitting low defaults to 0
	slice1 := originalSlice[:2] // ["A", "B"]
	fmt.Println("slice1:", slice1)
	// Omitting high defaults to len(slice)
	slice2 := originalSlice[3:] // ["D", "E"]
	fmt.Println("slice2:", slice2)

	//s := "()[]{}"
	s := "([])"
	fmt.Println("isValid:", isValid(s))
}

func isValid(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case ')', ']', '}':
			fmt.Println("current char:", string(char))
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return false
			}
			fmt.Println("length:", len(stack))
			fmt.Println("stack before pop:", string(stack))
			stack = stack[:len(stack)-1]
			fmt.Println("stack after pop:", string(stack))
		case '(', '[', '{':
			fmt.Println("appending char to stack:", string(char))
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

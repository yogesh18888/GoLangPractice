package main

import "fmt"

func main() {
	// Basic interpreted string
	s1 := "Hello, Go!"
	fmt.Println("s1:", s1)

	// String with escape sequences
	s2 := "This is a line.\nThis is another line.\tAnd a tab."
	fmt.Println("s2:\n", s2)

	// String with a double quote
	s3 := "He said, \"Hello!\""
	fmt.Println("s3:", s3)

	// Unicode escape sequence (represents 'A')
	s4 := "\u0041BC" // Unicode code point for 'A'
	fmt.Println("s4:", s4)

	// Hexadecimal byte escape sequence (represents 'A')
	s5 := "\x41BC"
	fmt.Println("s5:", s5)

	// Multi-byte Unicode characters (UTF-8 encoding)
	s6 := "Привет, мир! 👋" // Russian "Hello, world!" and a waving hand emoji
	fmt.Println("s6:", s6)
	fmt.Println("Length of s6 (bytes):", len(s6)) // len() returns byte count, not character count
}

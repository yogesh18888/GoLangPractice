package main

import "fmt"

// Custom error type for file operations
type FileError struct {
	Filename string
	Op       string // Operation: "open", "read", "write"
	Err      error  // The underlying error
}

func (e FileError) Error() string {
	return fmt.Sprintf("file error: failed to %s %s: %v", e.Op, e.Filename, e.Err)
}

// Example function returning a custom error
func openFile(filename string) (string, error) {
	if filename == "non_existent.txt" {
		// Simulate a file not found error
		return "", FileError{
			Filename: filename,
			Op:       "open",
			Err:      fmt.Errorf("file not found on disk"), // Underlying error
		}
	}
	return "file_content", nil
}

func main() {
	_, err := openFile("non_existent.txt")
	if err != nil {
		fmt.Println(err) // Output: file error: failed to open non_existent.txt: file not found on disk
		// You can use a type assertion to check if it's your custom error
		if fe, ok := err.(*FileError); ok {
			fmt.Printf("Specific FileError: Filename=%s, Operation=%s\n", fe.Filename, fe.Op)
		}
	}
}

package controller

import "fmt"

// Test handler function
func TestHandler(s string) string {
	return fmt.Sprintf("Test controller method %v\n", s)
}

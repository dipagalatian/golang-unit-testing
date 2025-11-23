package helper

import "testing"

// Unit test for HelloWorld function
// For testing, we can run golang command:
// - go test ./... (to run all tests in all subdirectories)
// - go test
// - go test -v (for verbose)
// - go test -run TestHelloWorldAdmin (to run specific test function)
// - go test -cover (to see code coverage)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Dipa")

	if result != "Hello Dipa!" {
		// Errror test
		t.Errorf("Expected 'Hello Dipa!' but got '%s'", result)
		panic("Test HelloWorld failed. Test did not return 'Hello Dipa!'")
	}
}

func TestHelloWorldAdmin(t *testing.T) {
	result := HelloWorld("Admin")

	if result != "Hello Admin!" {
		t.Errorf("Expected 'Hello Admin!' but got '%s'", result)
		panic("Test HelloWorld failed. Test did not return 'Hello Admin!'")
	}
}
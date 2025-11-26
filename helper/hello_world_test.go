package helper

import (
	"fmt"
	"testing"
)

// Unit test for HelloWorld function
// For testing, we can run golang command:
// - go test ./... (to run all tests in all subdirectories)
// - go test
// - go test -v (for verbose)
// - go test -run TestHelloWorldAdmin (to run specific test function)
// - go test -cover (to see code coverage)

// use *testing.T struct for mark test as failed
// use t.Fail() -> mark teset as failed but continue running rest code in the same test function
// use t.FailNow() -> mark test as failed and stop running rest code in the same test function
// use t.Errorf("format", args...) -> mark test as failed with formatted error message but continue running rest code in the same test function
// use t.Fatalf("format", args...) -> mark test as failed with formatted error message and stop running rest code in the same test function

func TestHelloWorldUser(t *testing.T) {
	result := HelloWorld("Dipa")

	if result != "Hello Dipa!" {
		// Errror test
		// t.Errorf("Expected 'Hello Dipa!' but got '%s'", result)
		// panic("Test HelloWorld failed. Test did not return 'Hello Dipa!'") -> this not recommended in testing, because it will creash/stop the all test process
		// instead use t.Fail() to mark the test as failed but continue running the rest code in the same test function
		t.Fail()
	}
	// this line will be executed even if the test fails above
	fmt.Println("TestHelloWorldUser done")
}

func TestHelloWorldAdmin(t *testing.T) {
	result := HelloWorld("Admin")

	if result != "Hello Admin!" {
		// t.Errorf("Expected 'Hello Admin!' but got '%s'", result)
		// panic("Test HelloWorld failed. Test did not return 'Hello Admin!'")
		t.FailNow() // this will stop the test process immediately
	}
	// this line will not be executed if the test fails above
	fmt.Println("TestHelloWorldAdmin done")
}
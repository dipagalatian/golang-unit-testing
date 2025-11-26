package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
// use t.Errorf("format", args...) -> mark test as failed with formatted error message but continue call t.Fail() running rest code in the same test function
// use t.Fatalf("format", args...) -> mark test as failed with formatted error message and call t.FailNow() stop running rest code in the same test function
// THE COMMMON PRACTICE IS TO USE PACKAGE "TESTIFY/ASSERT" FOR TESTING

// TestHelloWorldAssert uses testify/require package for testing
// if failed, will call t.FailNow()
// not return boolean value
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Dipa")

	require.Equal(t, "Hello Dipa!", result, "they should be equal")
	fmt.Println("TestHelloWorldRequire done")
}

// TestHelloWorldAssert uses testify/assert package for testing
// if failed, will call t.Fail()
// return boolean value
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Dipa")

	ar := assert.Equal(t, "Hello Dipa!", result, "they should be equal")
	fmt.Println("Assertion result:", ar)
	fmt.Println("TestHelloWorldAssert done")
}

func TestHelloWorldUser(t *testing.T) {
	result := HelloWorld("Dipa")

	if result != "Hello Dipa!" {
		// Errror test

		// Bad practice to use panic in testing
		// t.Errorf("Expected 'Hello Dipa!' but got '%s'", result)
		// panic("Test HelloWorld failed. Test did not return 'Hello Dipa!'") -> this not recommended in testing, because it will creash/stop the all test process

		// Good practice but not handling error message logs
		// instead use t.Fail() to mark the test as failed but continue running the rest code in the same test function
		// t.Fail()

		// Best practice to use t.Error() -> logs error + t.Fail()
		t.Error("Expected 'Hello Dipa!' but got '" + result + "'")
	}
	// this line will be executed even if the test fails above
	fmt.Println("TestHelloWorldUser done")
}

func TestHelloWorldAdmin(t *testing.T) {
	result := HelloWorld("Admin")

	if result != "Hello Admin!" {
		// Error test

		// Bad practice to use panic in testing
		// t.Errorf("Expected 'Hello Admin!' but got '%s'", result)
		// panic("Test HelloWorld failed. Test did not return 'Hello Admin!'")

		// Good practice but not handling error message logs
		// t.FailNow() // this will stop the test process immediately

		// Best practice to use t.Fatal() -> logs error + t.FailNow()
		t.Fatal("Expected 'Hello Admin!' but got '" + result + "'")
	}
	// this line will not be executed if the test fails above
	fmt.Println("TestHelloWorldAdmin done")
}
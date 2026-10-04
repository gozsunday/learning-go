package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}

// t.Errorf can be used for a Printf-style formatting string

func Test_subtractNumbers(t *testing.T) {
	result := subtractNumbers(5, 2)
	if result != 3 {
		t.Errorf("incorrect result: expected 3, got %d", result)
	}
}

// when t.Error and t.Errorf are used, the test function continues processing.
// if it is desired to stop the test once an error occurs, use the t.Fatal or t.Fatalf
// methods.

// ——————————————————————————————
// Setting up and tearing down
// ——————————————————————————————

// there are cases when a common state has to be set up before any tests run, and then
// removed after testing is complete to manage this state and also run the tests, a
// TestMain function is used.

// once there's a TestMain in a package, `go test` runs it instead of running the
// individual tests.
// it then becomes the job of the TestMain function to set up and run the tests.

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("Set up stuff for tests here")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Clean up stuff after tests here")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("TestFirst uses stuff set up in TestMain", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("TestSecond also uses stuff set up in TestMain", testTime)
}

// TestMain is useful in cases where you need to set up data from say, an external database,
// or cases where the code being tested depends on package-level variables.

// t.CleanUp

// t.CleanUp is called to clean up temp resources created for a single test.
// it has a single parameter, a function with no input parameters or return values.
// this function runs when the tests complete.
// for simple tests, defer can be used.
// t.CleanUp can be used multiple times in a single test and it runs in a LIFO order.

// helper function called from multiple tests showing t.CleanUp in use.
func createFileOne(t *testing.T) (_ string, err error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()
	// write some data to f
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func TestFileProcessingOne(t *testing.T) {
	_, err := createFileOne(t)
	if err != nil {
		t.Fatal(err)
	}
	// do testing, dont worry about cleanup
}

// if a test uses temp files, we can avoid writing cleanup code by using the TempDir
// method on *testing.T. it creates a new temp directory everytime it is invoked and
// returns the full path of the directory. it also registers a handler with Cleanup to
// delete the directory and its contents when the test has completed.

func createFileTwo(tempDir string) (_ string, err error) {
	f, err := os.CreateTemp(tempDir, "tempFile")
	if err != nil {
		return "", err
	}
	// no t.CleanUp since os.CreateTemp does its own cleanup.
	defer func() {
		err = errors.Join(err, f.Close())
	}()
	return f.Name(), nil
}

func TestFileProcessingTwo(t *testing.T) {
	tempDir := t.TempDir()
	_, err := createFileTwo(tempDir)
	if err != nil {
		t.Fatal(err)
	}
	// do testing, dont worry about cleanup
}

// ——————————————————————————————
// Testing with environment variables
// ——————————————————————————————

// go provides a helper method on *testing.T for to register env vars.

type EnvVars struct {
	Port string
}

func ProcessEnvVars() EnvVars {
	return EnvVars{
		Port: "8080",
	}
}

func TestEnvVarProcess(t *testing.T) {
	t.Setenv("PORT", "8080")
	cfg := ProcessEnvVars()
	if cfg.Port == "" {
		t.Error("Port not set in env vars")
	}
}

// ——————————————————————————————
// Storing sample test data
// ——————————————————————————————

// when `go test` is ran, it uses the package directory of any test being run as the
// current working dir. to have sample data for testing in a package, create a
// subdirectory called testdata to hold the files. this is a reserved dir name.

// when reading from testdata, always use a relative file path.

// ——————————————————————————————
// Testing your public API
// ——————————————————————————————

// tests need to be placed in the same package as the functions they're testing.
// this means they can also access private functions. to ensure that the tests can
// only access public functions, go provides a way to do that. appending a `_test`
// to the package name i.e. renaming the package name in the test file as
// `packagename_test`.

// ——————————————————————————————
// Using go-cmp
// ——————————————————————————————

// writing a thorough comparison of a compound type's two instances can be verbose.
// we can use google's go-cmp to do this.

func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name:      "Goziem",
		Age:       16,
		DateAdded: time.Now(),
	}
	result := CreatePerson("Goziem", 16)

	// this test will fail because the expexted DateAdded wont match with the result
	// DateAdded, and this can cause issues whrn testing code.

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}

	// to ignore the DateAdded,
	// we can use cmp.Comparer, then use it in cmp.Diff

	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})
	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
}

// ——————————————————————————————
// Running Table Tests
// ——————————————————————————————

// table tests are used when multiple test cases need to be used to validate that a
// function is working well. instead of writing a lot of repetitive testing logic,
// we can use table tests.

// lets try to test the function DoMath()

func TestDoMath(t *testing.T) {
	result, err := DoMath(8, 2, "+")
	if result != 10 {
		t.Errorf("should have been 10; got %d", result)
	}
	if err != nil {
		t.Errorf("should have been nil error; got %v", err)
	}

	result, err = DoMath(8, 2, "-")
	if result != 6 {
		t.Errorf("should have been 6; got %d", result)
	}
	if err != nil {
		t.Errorf("should have been nil error; got %v", err)
	}

	result, err = DoMath(8, 2, "*")
	if result != 16 {
		t.Errorf("should have been 16; got %d", result)
	}
	if err != nil {
		t.Errorf("should have been nil error; got %v", err)
	}

	result, err = DoMath(8, 2, "/")
	if result != 4 {
		t.Errorf("should have been 4; got %d", result)
	}
	if err != nil {
		t.Errorf("should have been nil error; got %v", err)
	}
}

// as you can see, the above test is quite repetitive
// let's replace this repetition with a table test

func TestDoMathTable(t *testing.T) {
	// create a slice of anon structs
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"addition", 8, 2, "+", 10, ""},
		{"subtraction", 8, 2, "-", 6, ""},
		{"multiplication", 8, 2, "*", 16, ""},
		{"division", 8, 2, "/", 4, ""},
		{"bad_division", 8, 0, "/", 0, "division by zero"},
	}

	// loop through the data and run the test on each case
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}

			// comapring error msgs can be tricky, especially when

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error message '%s', got '%s'", d.errMsg, errMsg)
			}
		})
	}
}

// ——————————————————————————————
// Running Tests Concurrently
// ——————————————————————————————

// by default, unit tests run independently from every other unit test.
// this makes them good candidates for concurrency.
// to make them run in parallel, we can add `t.Parallel()` as the first line of the test.

// parallel tests can speed up large, long-running test suites.
// they do have some disadvantages tho. it isnt advised to mark tests that touch the
// same shared mutable state as parallel, because the results can be inconsistent.

// ——————————————————————————————
// Checking Code Coverage
// ——————————————————————————————

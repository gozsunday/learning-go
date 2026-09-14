package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
	"slices"
)

func main() {
	Errors()
	SentinelErrors()
	WrappedErrors()
	IsAndAs()
}

func Errors() {
	//* Errors
	// Go handles errors by returning a value of type error as the last return value of a function.
	// When a function executes as expected, `nil` is returned as the error.
	// If something goes wrong, an error value is returned. The calling function then checks the
	// error return value by comparing it with `nil`, handling the error or returning an error of
	// its own.

	// error msgs should not be capitalized nor should they end with a punctuation or new line.
	// when an error value is returned, the other return values should be set to their zero values.

	numerator := 20
	denominator := 3

	remainder, mod, err := calcRemainderAndMod(numerator, denominator)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(remainder, mod)

	// there are two ways to create an error from a string:
	// `errors.New("string")`: takes in a string and returns an error
	// `fmt.Errorf("%v string", i)`: this allows to include runtime values in the error msg
}

func calcRemainderAndMod(num, den int) (int, int, error) {
	if den == 0 {
		return 0, 0, errors.New("denominator is 0")
	}

	return num / den, num % den, nil
}

func SentinelErrors() {
	//* Sentinel Errors
	// Some errors are meant to signify that processing cannot continue because of a problem
	// with the current state. We call these Sentinel Errors.

	// Sentinel errors are one of the few variables that are declared at the package level.
	// They start with `Err` with the exception of io.EOF

	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("told you so")
	}
}

func ErrorsAreValues() {
	// since `error` is an interface, we can define our own errors that include additional info
	// for logging or error handling.

	//? check `LoginAndGetData("", "", "")`

	// if youre using your own error type, be sure to not return an uninitialized instance.
	// You can return the error (or nil) instead of creating a var with the custom error type
	// you can also assign the inbuilt `error` type to the error you want to return even if you
	// plan on using a custom error
}

func WrappedErrors() {
	//* WRAPPED ERORS
	// When an error is passed through the code, we might want to add some info or context to it.
	// this can be done by wrapping the error using fmt.Errorf(). The verb used to embed the
	// error is `%w`.
	// the wrapped error can also be unwrapped using errors.Unwrap().

	err := FileChecker("not_real.txt")
	if err != nil {
		fmt.Println(err)
		if wErr := errors.Unwrap(err); wErr != nil {
			fmt.Println(wErr)
		}
	}

	// Apparently, we dont usually call errors.Unwrap() directly. errors.Is() and errors.As() is
	// usually used. More on this later.

	// If we want to wrap an error with our own custom error, it has to implement the `Unwrap` method.
	// The `Unwrap` method takes in no parameters and returns an `error`.
	//? Check `StatusErr`

	// if we want to create an eror that contains the message from another error without wrapping it, we use the `%v` verb instead of `%w`.

	//* WRAPPING MULTIPLE ERRORS
	// Sometimes a function generates multiple errors that should be returned.
	// example, imagine a function for validating the fields of a struct, it would be better to
	// return an error for each invalid field. since the standard function signature returns
	// `error` instead of `[]error`, we need to merge multiple errors into one. we use errors.Join()
	// for this.

	person1 := Person{
		FirstName: "",
		LastName:  "",
		Age:       -21,
	}

	err = ValidatePerson(person1)
	fmt.Println(err)

	// another way to merge multiple errors is to pass multiple `%w` verbs to `fmt.Errorf`

	err1 := errors.New("first error")
	err2 := errors.New("second error")
	err3 := errors.New("third error")
	err = fmt.Errorf("first: %w\nsecond: %w\nthird: %w", err1, err2, err3)

	fmt.Println(err)

	// We can implement a custom error type that supports multiple wrapped errors, by adding the
	// `Unwrap()` method to the custom error and have it return `[]error`.

	//? Goto MyError

	// Note that, passing an error that implements the `[]error` variant of `Unwrap()`to the
	// `errors.Unwrap()` function will return `nil`.
}

func IsAndAs() {
	//* Is
	// wrapping errors is really useful for getting additional info about an error, but it has issues.
	// if a Sentinel error is wrapped, there's no way to check for it using "==".
	// a type assertion or type switch can't be used to match a wrapped custom error.

	// Go solves these issues using `errors.Is()` and `errors.As()`.
	// `errors.Is()` takes in two parameters: the error being checked, and the instance we're
	// comparing it against. It returns true if any error in the error tree matches the
	// provided sentinel error.

	// working with the `FileChecker(name string) error`, we have:

	err := FileChecker("fake.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}

	// To implement this in our own custom errors, we need to define a Is() method.
	//? Check MyError.Is()

	//* As
	// the `errors.As()` allows us to check whether a returned error (or any error it wraps) matches
	// a specific type. it takes in two parameters: the error being examined, and a pointer to a
	// variable of the type that you're looking for.

	// if the function returns true, an error in the error tree was found that matched and the
	// matching error is assigned to the second parameter. if the func returns false, no match
	// was found in the error tree.

	// trying with the `FileChecker(name string) error`, we have:

	err = FileChecker("no_text.txt")
	var myErr MyErr
	if errors.As(err, &myErr) {
		fmt.Println(myErr.Codes)
	}
}

type Status int

const (
	Unauthorized Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
	Err     error //? for wrapping errors
}

func (se StatusErr) Error() string {
	return se.Message
}

// ? for wrapping errors
func (se StatusErr) Unwrap() error {
	return se.Err
}

func Login(uid, pwd string) (string, error) {
	var token string
	return token, nil
}
func GetData(token, file string) ([]byte, error) {
	var data []byte
	return data, nil
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	token, err := Login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  Unauthorized,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
			Err:     err, //? for wrapping errors
		}
	}

	data, err := GetData(token, file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
			Err:     err, //? for wrapping errors
		}
	}

	return data, nil
}

func FileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in FileChecker: %w", err)
	}

	f.Close()
	return nil
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func ValidatePerson(p Person) error {
	var errs []error

	if len(p.FirstName) == 0 {
		errs = append(errs, errors.New("field 'FirstName' cannot be empty"))
	}
	if len(p.LastName) == 0 {
		errs = append(errs, errors.New("field 'LastName' cannot be empty"))
	}
	if p.Age < 0 {
		errs = append(errs, errors.New("field 'Age' cannot be negative"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

type MyError struct {
	Code   int
	Errors []error
}

func (me MyError) Error() string {
	return errors.Join(me.Errors...).Error()
}

func (me MyError) Unwrap() []error {
	return me.Errors
}

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

// ? for IsAndAs()
func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return slices.Equal(me.Codes, me2.Codes)
	}
	return false
}

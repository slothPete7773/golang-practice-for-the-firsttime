# Error Handling

Go does not have Try...Catch, nor Exception Class

We have to manually implement Error/Exception Handling.

In Go, error handling is handled via the Error return from the function.

Note 
- By default, Go's function returns 2 things; `func() -> [Value], Error`.
  - If no error occur, `Error` will be `nil` or null.
  - Otherwise, it will be something else in case of error.
- Go has `errors` package.

**Example**
```go
import (
	"errors"
	"fmt"
)
// divide divides two integers and returns an error if the divisor is 0
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}
```

From the snippet above:
1. The function `divide()` also return `errors.New("cannot divide by zero")` from package `errors` to indicate the error.
2. And the expected result in `main()` from function `divide(10, 0)` also expect actual `result` and possible `err`.


## Custom Error

We can also create a custom Error using Go's Interface and Struct.

```go
type error interface {
  Error() string
}
```

The following is an example for Login Error

1. We declare a LoginError struct for use as Error object and an interface for that LoginError struct to behave when error happened.
2. Then, create a function `Error()` with a binding-receiver argument for LoginError struct.
3. And a function for simulating the log in.
```go

type error interface {
  Error() string
}

// LoginError is a custom error type for login failures
type LoginError struct {
	Username string
	Message  string
}

// Implement the Error() method to satisfy the error interface
func (e *LoginError) Error() string {
	return fmt.Sprintf("Login error for user '%s': %s", e.Username, e.Message)
}

// Simulated function that attempts a user login
func login(username, password string) error {
	// Simulate checking username and password
	if username != "admin" || password != "password123" {
		return &LoginError{Username: username, Message: "invalid credentials"}
	}
	// Login successful
	return nil
}

func main() {
	// Attempt to login with wrong credentials
	err := login("user", "pass")
	if err != nil {
		switch e := err.(type) {
		case *LoginError:
			// Custom error handling
			fmt.Println("Custom error occurred:", e)
		default:
			// Other types of errors
			fmt.Println("Generic error occurred:", e)
		}
		return
	}

	// Continue with the rest of the program if login is successful
	fmt.Println("Login successful!")
}

```
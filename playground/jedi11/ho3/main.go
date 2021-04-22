package main

import "fmt"

type customError struct {
	message string
}

func (ce customError) Error() string {
	return fmt.Sprintf("big time error: %v", ce.message)
}

func main() {
	custErr := customError{
		message: "There has been an error, abort the prog",
	}
	foo(custErr)
}

func foo(err error) {
	fmt.Println("Reached funcntion foo:", err, "\n", err.(customError).message)
}

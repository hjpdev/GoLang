package main // In Go, it's idiomatic to communicate errors via an explicit, seperate return value

import "errors"
import "fmt"

func f1(arg int) (int, error) { // By convention, errors are the last return value, and have type error - a built in interface
	if arg == 42 {
		return -1, errors.New("Can't work with 42") // This constructs a basic error, with the givwn message
	}
	return arg + 3, nil // A nil value implies there was no error
}

type argError struct { // It's possible to use custom error types by implementing the Error() method on them
	arg  int // This variant uses a custom type to explicitly represent an argument error
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 { // Here &argError is used to build a new struct
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() { // The two loops below test out each of the error-returning functions
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil { // Use of an inline error check on the if line is a common idiom in Go code
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, e := f2(42) // If you want to programmatically use the data in a custom error, you’ll need to get the error as an instance of the custom error type via type assertion
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

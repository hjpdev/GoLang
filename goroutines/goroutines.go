package main // A Goroutine is a lightweight thread of execution

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct") // This runs the function in the normal way - synchronously

	go f("goroutine") // To invoke the function in a Goroutine, it is pprefixed with 'go'
	// This new Goroutine will execute concurrently with the calling one
	go func(msg string) { // Can also start a Goroutine for an anonymous function call
		fmt.Println(msg)
	}("going") // Our two function calls are running asynchronously in separate goroutines now, so execution falls through to here

	fmt.Scanln() // Scanln requires we press a key before the program exits
	fmt.Println("done")
}

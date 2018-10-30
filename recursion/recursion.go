package main

import "fmt"

func fact(n int) int {
	if n == 0 { // This is the base case
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
	fmt.Println(fact(4))
	fmt.Println(fact(5))
	fmt.Println(fact(3))
}

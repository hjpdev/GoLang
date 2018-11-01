package main

import "fmt"

// Calculate ...
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

// Sum ...
func Sum(x int, y int) (result int) {
	result = x + y
	return result
}

func main() {
	result1 := Calculate(2)
	result2 := Sum(1, 2)

	fmt.Println(result1)
	fmt.Println(result2)
}

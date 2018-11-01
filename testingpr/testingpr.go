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

// Power ...
func Power(x int, y int) (result int) {
	result = 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}

func main() {
	result1 := Calculate(2)
	result2 := Sum(1, 2)
	result3 := Power(2, 3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

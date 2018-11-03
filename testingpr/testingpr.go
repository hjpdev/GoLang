package main

import "fmt"

// Plus2 ...
func Plus2(x int) (result int) {
	result = x + 2
	return result
}

// Sum ...
func Sum(x int, y int) (result int) {
	result = x + y
	return result
}

// Minus ...
func Minus(x int, y int) (result int) {
	result = x - y
	return result
}

// Divider ...
func Divider(x float32, y float32) (result float32) {
	result = x / y
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
	result1 := Plus2(2)
	result2 := Sum(1, 2)
	result3 := Power(2, 3)
	result4 := Minus(6, 4)
	result5 := Divider(8, 2)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
}

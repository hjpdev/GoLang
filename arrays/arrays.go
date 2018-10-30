package main

import "fmt"

func main() {

	var a [5]int           // This creates an array which will hold 5 ints
	fmt.Println("emp:", a) // By default an array is zero-valued, which for ints means 0

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) // Returns length

	b := [5]int{1, 2, 3, 4, 5} // Use this syntax to declare & initialize an array in the same line
	fmt.Println("dcl", b)

	var twoD [2][3]int // For multi-dimensional arrays
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

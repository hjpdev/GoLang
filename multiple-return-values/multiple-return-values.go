package main

import "fmt"

func vals() (int, int) { // (int, int) shows it returns 2 ints
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // Return subset with blank operator _
	fmt.Println(c)
}

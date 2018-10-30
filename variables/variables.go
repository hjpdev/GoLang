package main

import "fmt"

func main() {

	var a = "initial" // var declares one or more variables
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true // Go will infer the type of initialised variables
	fmt.Println(d)

	var e int      // Variables declared without a corresponding initialistion are zero-valued
	fmt.Println(e) // e.g. The zero-value for an int is 0

	f := "short" // The := syntax is shorthand for declaring & initialising a variable
	fmt.Println(f)
}

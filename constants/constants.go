package main

import "fmt"
import "math"

const s string = "constant" // A const statement can appear anywhere a var statement can

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n // Constant expressions perform arithmetic with arbitray precision
	fmt.Println(d)

	fmt.Println(int64(d)) // A numeric constant has no type until given pne, such as by an explicit cast

	fmt.Println(math.Sin(n)) // A number can be given a type by using it in a context taht requires one e.g. here Sin expects a float64
}

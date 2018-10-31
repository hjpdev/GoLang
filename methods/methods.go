package main // Go supports methods defined on struct types

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int { // Methods can be defined for either pointer or value receiver types
	return 2*r.width + 2*r.height // This is an example of a value receiver
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())
}

package main // Interfaces are named collections of method signatures

import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) { // If a variable has an interface type, then we can call methods that are in the named interface
	fmt.Println(g) // Here’s a generic measure function taking advantage of this to work on any geometry.
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// circle & rect both implement geometry interface ∴ can use instances of these structs as arguments to measure

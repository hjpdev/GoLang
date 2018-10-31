package main // Go's structs are collections of fields, useful for grouping data together to form records

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"}) // Omitted fields will be zero-valued

	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // Access struct fields with a dot

	sp := &s
	fmt.Println(sp.age) // You can also use dots with struct pointers - the pointers are automatically dereferenced

	sp.age = 51 // Structs are mutable
	fmt.Println(sp.age)
}

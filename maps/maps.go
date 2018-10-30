package main

import "fmt"

func main() { // Maps are Go’s built-in associative data type, sometimes called hashes or dicts in other languages

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m) // Printing a map with e.g. fmt.Println will show all of its key/value pairs.

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs) // The optional second return value indicates if the key was present in the map
	// This can be used to disambiguate between missing keys and keys with zero values
	n := map[string]int{"foo": 1, "bar": 2} // Here we didn’t need the value itself, so we ignored it with the blank identifier _
	fmt.Println("map:", n)
}

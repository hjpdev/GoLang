package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums { // range on arrays and slices provides both the index and value for each entry.
		if num == 3 { // Above we didn’t need the index, so we ignored it with the blank identifier _
			fmt.Println("index", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"} // range on map iterates over key/value pairs
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs { // range can also iterate over just the keys of a map
		fmt.Println("key", k)
	}
	for i, c := range "go" { // range on strings iterates over Unicode code points.
		fmt.Println(i, c) // The first value is the starting byte index of the rune and the second the rune itself
	}
}

package main

import "fmt"

func main() { // Unlike arrays, slices are typed only by the elements they contain, not no. of elements

	s := make([]string, 3) // This makes a slice of strings of length 3
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:,", c)

	l := s[2:5] // Slices support a slice operator, where the syntax is Slice[low:high]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3) // Slices can be composed into multi-dimensional data structures.
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j // The length of the inner slices can vary, unlike with multi-dimensional arrays.
		}
	}
	fmt.Println("2d:", twoD)
}

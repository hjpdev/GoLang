package main // Pointers allow you to pass references to values/records

import "fmt"

func zeroval(ival int) {
	ival = 0 // zeroval will get a copy ival, distinct from the one in the calling function
}

func zeroptr(iptr *int) {
	*iptr = 0 // The *iptr code here dereferences the pointer from its memory address to the current value at that address
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // the &i syntax gives the memory address of i, i.e. a pointer to i
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i) // Pointers can be printed too
}

/*
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0xc000016188
*/

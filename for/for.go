package main

import "fmt"

func main() {
	i := 1 // For is Go's only looping construct, and there are 3 basic implementations

	for i <= 3 { // The most basic with a single condition
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ { // A classic initial/condition/after for loop
		fmt.Println(j)
	}

	for { // for without a condition will loop repeatedly until you break/return out from the enclosing function
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ { // You can also continue to the next iteration of the loop
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

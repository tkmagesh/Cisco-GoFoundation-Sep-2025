package main

import "fmt"

func main() {
	var x int
	x = 100

	// referencing (value -> address)
	var xPtr *int
	xPtr = &x
	fmt.Println("Address of x :", xPtr)

	// dereferencing (address -> value)
	xVal := *xPtr
	fmt.Println("xVal :", xVal)

	fmt.Println("x == *(&x) :", x == (*(&x)))
}

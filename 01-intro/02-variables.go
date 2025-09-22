package main

import "fmt"

// package scope allows un-used variables
var x int

// cannot use ":=" in package
// y := 200

func main() {

	/*
		var username string
		username = "Magesh"
	*/

	/*
		var username string = "Magesh"
	*/

	// type inference (only when declaration & initialization done together)
	/*
		var username = "Magesh"
	*/

	// using ":=" (only when declaration & initialization done together)
	username := "Magesh"
	fmt.Printf("Hello %s, Have a nice day!\n", username)

	// multiple variables

	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	// combine the declaration of variables of the same type together
	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	// combine the declaration of variables together irrespective of their type
	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	// declaration & initialization together
	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of %d and %d is %d\n"
			result int    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// type inference
	/*
		var (
			x, y   = 100, 200
			str    = "Sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// Using ":="

	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)

}

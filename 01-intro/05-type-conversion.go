package main

import "fmt"

func main() {
	var x int8 = 100
	var f float64

	// use the typename like a function for type conversion
	f = float64(x)

	fmt.Println(f)
}

package main

import "fmt"

func main() {
	var x int
	x = 100
	fmt.Println("[main] &x =", &x)
	fmt.Printf("Before incrementing, x = %d\n", x)
	increment(&x)
	fmt.Printf("After incrementing, x = %d\n", x)
}

func increment(no *int) /* no return values */ {
	// fmt.Println("[increment] &no =", &no)
	*no += 1
}

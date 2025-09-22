package main

import "fmt"

func main() {

	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
		fmt.Println("no =", no)
	*/

	// "no" variable is scoped to the if-else block
	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	// "no" variable is not accessible
	//fmt.Println("no =", no)

}

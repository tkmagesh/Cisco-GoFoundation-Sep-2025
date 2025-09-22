package main

import "fmt"

func main() {
	/*
		var username string
		fmt.Println("Enter your name :")
		// username = "Magesh"
		fmt.Scanln(&username)
		fmt.Printf("Hello %s, Have a nice day!\n", username)
	*/

	var n1, n2 int
	fmt.Println("Enter 2 numbers :")
	fmt.Scanln(&n1, &n2)
	fmt.Printf("n1 = %d and n2 = %d\n", n1, n2)
}

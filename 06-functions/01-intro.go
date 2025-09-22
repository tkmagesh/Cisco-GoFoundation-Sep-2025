package main

import "fmt"

func main() {
	sayHi()
	sayHello("Magesh")
	greet("Magesh", "Kuppan")
	fmt.Println(getGreet("Suresh", "Kannan"))

	// fmt.Println(divide(100, 7))

	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	// print only quotient
	/* _ is used to RECIEVE a value that we dont want to use  */
	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
	*/
}

func sayHi() {
	fmt.Println("Hi there!")
}

func sayHello(userName string) {
	fmt.Printf("Hello %s!\n", userName)
}

/*
func greet(firstName string, lastName string) {
	fmt.Printf("Hi %s %s! Have a nice day!\n", firstName, lastName)
}
*/

func greet(firstName, lastName string) {
	fmt.Printf("Hi %s %s! Have a nice day!\n", firstName, lastName)
}

func getGreet(firstName, lastName string) string {
	return fmt.Sprintf("Hi %s %s! Have a good day!", firstName, lastName)
}

/*
func divide(multiplier, divisor int) (int, int) {
	quotient := multiplier / divisor
	remainder := multiplier % divisor
	return quotient, remainder
}
*/

// Named results
func divide(multiplier, divisor int) (quotient, remainder int) {
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return
}

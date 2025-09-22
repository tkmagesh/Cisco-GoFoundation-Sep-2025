package main

import "fmt"

func main() {

	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()

	var sayHello func(string)
	sayHello = func(userName string) {
		fmt.Printf("Hello %s!\n", userName)
	}
	sayHello("Magesh")

	var sum func(...int) int
	sum = func(list ...int) int {
		var result int
		for idx := 0; idx < len(list); idx++ {
			result += list[idx]
		}
		return result
	}
	total := sum(10, 20, 30, 40)
	fmt.Println("total :", total)

	// var divide func(int, int) (int, int)
	var divide func(m, d int) (q, r int)

	divide = func(multiplier, divisor int) (quotient, remainder int) {
		quotient = multiplier / divisor
		remainder = multiplier % divisor
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}

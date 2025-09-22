package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hi there!")
	}()

	func(userName string) {
		fmt.Printf("Hello %s!\n", userName)
	}("Magesh")

	total := func(list ...int) int {
		var result int
		for idx := 0; idx < len(list); idx++ {
			result += list[idx]
		}
		return result
	}(10, 20, 30, 40)
	fmt.Println("total :", total)

	q, r := func(multiplier, divisor int) (quotient, remainder int) {
		quotient = multiplier / divisor
		remainder = multiplier % divisor
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}

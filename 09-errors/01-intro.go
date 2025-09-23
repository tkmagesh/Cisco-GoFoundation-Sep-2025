package main

import (
	"errors"
	"fmt"
)

func main() {
	// divisor := 7
	divisor := 0
	q, r, err := divide(100, divisor)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

/*
func divide(multiplier, divisor int) (int, int, error) {
	var quotient, remainder int
	if divisor == 0 {
		e := errors.New("divisor cannot be 0")
		return 0, 0, e
	}
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return quotient, remainder, nil
}
*/

func divide(multiplier, divisor int) (quotient, remainder int, e error) {
	if divisor == 0 {
		e = errors.New("divisor cannot be 0")
		return
	}
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return
}

package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {
	// divisor := 7
	divisor := 0
	q, r, err := divide(100, divisor)
	if err == ErrDivideByZero {
		fmt.Println("Do not attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println("Something went wrong, Error :", err)
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divide(multiplier, divisor int) (quotient, remainder int, e error) {
	if divisor == 0 {
		e = ErrDivideByZero
		return
	}
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return
}

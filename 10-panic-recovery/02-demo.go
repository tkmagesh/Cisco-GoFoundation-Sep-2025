package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("Divide by 0 error")

func main() {

	var divisor int

	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		q, r, err := divideAdapter(100, divisor)
		if err != nil {
			fmt.Println("Error occurred, try a new divisor : ", err)
			continue
		}
		fmt.Printf("[main] Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
		break
	}
}

// converts the panic from 'divide' function into an error
func divideAdapter(multiplier, divisor int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(multiplier, divisor)
	return
}

// 3rd party api (WE CANNOT change the code)
func divide(multiplier, divisor int) (quotient, remainder int) {
	defer fmt.Println("	[divide] - deferred")
	fmt.Println("[divide] - calculating quotient")
	if divisor == 0 {
		panic(ErrDivideByZero)
	}
	quotient = multiplier / divisor
	fmt.Println("[divide] - calculating remainder")
	remainder = multiplier % divisor
	return
}

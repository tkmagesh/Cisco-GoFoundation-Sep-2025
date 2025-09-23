package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrDivideByZero error = errors.New("Divide by 0 error")

func main() {
	defer func() {
		fmt.Println("	[main] - deferred")
		if e := recover(); e != nil {
			log.Println("App panicked, Error : ", e)
			return
		}

	}()
	// divisor := 7
	divisor := 0
	q, r := divide(100, divisor)
	fmt.Printf("[main] Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

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

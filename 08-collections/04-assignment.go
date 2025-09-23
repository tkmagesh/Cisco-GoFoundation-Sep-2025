package main

import "fmt"

/*
Accept "start" and "end" from the user
Print all the prime numbers between "start" and "end" (including start and end)
*/

/* Refactor the solution using functions. Make sure that each function has ONLY ONE reason to change (SRP) */

func main() {
	var start, end int
	fmt.Println("Enter the start and end :")
	fmt.Scanln(&start, &end)
	primeNos := generatePrimes(start, end)
	for _, primeNo := range primeNos {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
}

func generatePrimes(start, end int) []int {
	var primes []int
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

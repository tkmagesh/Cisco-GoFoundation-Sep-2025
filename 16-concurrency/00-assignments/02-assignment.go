package assignments
package main

import (
	"fmt"
	"sync"
)

/* move the logic of print prime number from the "checkAndPrintPrime" to "main" function */

func main() {
	var start, end int
	wg := &sync.WaitGroup{}
	fmt.Println("Enter start and end:")
	fmt.Scanln(&start, &end)

	for no := start; no <= end; no++ {
		wg.Add(1)
		go checkAndPrintPrime(no, wg)
	}
	wg.Wait()
	fmt.Println("Done!")
}

func checkAndPrintPrime(no int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
}

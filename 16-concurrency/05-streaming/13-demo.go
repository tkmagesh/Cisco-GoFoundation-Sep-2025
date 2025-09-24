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
	ch := make(chan int)
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if IsPrime(no) {
				ch <- no
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for primeNo := range ch {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

func IsPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

package main

import (
	"fmt"
	"sync"
)

/* move the logic of print prime number from the "checkAndPrintPrime" to "main" function */

var primeNos []int
var mutex sync.Mutex

func main() {
	var start, end int
	wg := &sync.WaitGroup{}
	fmt.Println("Enter start and end:")
	fmt.Scanln(&start, &end)

	for no := start; no <= end; no++ {
		wg.Add(1)
		go checkPrime(no, wg)
	}
	wg.Wait()
	for _, primeNo := range primeNos {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

func checkPrime(no int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	mutex.Lock()
	{
		primeNos = append(primeNos, no)
	}
	mutex.Unlock()
}

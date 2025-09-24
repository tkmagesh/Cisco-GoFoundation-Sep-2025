package main

import "fmt"

/* process the numbers concurrently */

func main() {
	var start, end int
	fmt.Println("Enter start and end:")
	fmt.Scanln(&start, &end)
LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		fmt.Println("Prime No :", no)
	}
	fmt.Println("Done!")
}

package main

import "fmt"

func main() {
	fmt.Println(sum(10))             //=> 10
	fmt.Println(sum(10, 20))         //=>30
	fmt.Println(sum(10, 20, 30, 40)) //=> 100
	fmt.Println(sum())               //=> 0
}

func sum(list ...int) int {
	var result int
	for idx := 0; idx < len(list); idx++ {
		result += list[idx]
	}
	return result
}

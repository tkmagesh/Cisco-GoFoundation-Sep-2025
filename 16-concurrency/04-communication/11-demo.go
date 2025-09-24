package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- add(100, 200)
	}()
	result := <-ch
	fmt.Println(result)
}

func add(x, y int) int {
	time.Sleep(3 * time.Second)
	return x + y
}

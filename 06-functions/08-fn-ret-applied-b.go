package main

import (
	"fmt"
	"log"
)

func main() {
	add := wrapLogOperation(add)
	subtract := wrapLogOperation(subtract)

	addResult := add(100, 200)
	fmt.Println("Add Result :", addResult)

	subResult := subtract(100, 200)
	fmt.Println("Subtract Result :", subResult)

}

// ver 1.0
func add(x, y int) int {
	fmt.Printf("Processing [Add] x = %d and y = %d\n", x, y)
	return x + y
}

func subtract(x, y int) int {
	fmt.Printf("Processing [Subtract] x = %d and y = %d\n", x, y)
	return x - y
}

func wrapLogOperation(op func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		log.Println("Operation started...")
		result := op(x, y)
		log.Println("Operation completed!")
		return result
	}
}

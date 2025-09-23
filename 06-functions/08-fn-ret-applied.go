package main

import (
	"fmt"
	"log"
)

func main() {
	add := wrapLogOperation(add)
	subtract := wrapLogOperation(subtract)

	add(100, 200)
	subtract(100, 200)

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
		logOperation(func(x, y int) {
			fmt.Println("Multiply Result :", x*y)
		}, 100, 200)
	*/
	/*
		logAdd := wrapLogOperation(add)
		logAdd(100, 200)

		logSubtract := wrapLogOperation(subtract)
		logSubtract(100, 200)
	*/

}

// ver 1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

// ver 2.0
/*
func logAdd(x, y int) {
	log.Println("Operation started...")
	add(x, y)
	log.Println("Operation completed!")
}

func logSubtract(x, y int) {
	log.Println("Operation started...")
	subtract(x, y)
	log.Println("Operation completed!")
}
*/

// ver3.0
// Apply commonality-variability to the above
/*
func logOperation(op func(int, int), x, y int) {
	log.Println("Operation started...")
	op(x, y)
	log.Println("Operation completed!")
}
*/

// ver 4.0
// Apply "function composition"
func wrapLogOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started...")
		op(x, y)
		log.Println("Operation completed!")
	}
}

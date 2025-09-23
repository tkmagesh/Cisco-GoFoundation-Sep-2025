package main

import "fmt"

func main() {
	/*
		var productRanks map[string]int = make(map[string]int)
		productRanks["pen"] = 3
		productRanks["pencil"] = 1
		productRanks["marker"] = 5
	*/

	// var productRanks map[string]int = map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	// var productRanks = map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	// productRanks := map[string]int{"pen": 3, "pencil": 1, "marker": 5}

	productRanks := map[string]int{
		"pen":    3,
		"pencil": 1,
		"marker": 5,
	}
	fmt.Println(productRanks)
	fmt.Println("len(productRanks) : ", len(productRanks))

	fmt.Printf("Iterating a map [using range]\n")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// check for the existence of a key
	// keyToCheck := "pencil"
	keyToCheck := "scribble-pad" // non-existent key
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key %q does not exist\n", keyToCheck)
	}

	// remove an item
	// keyToRemove := "pen"
	keyToRemove := "scribble-pad" //non-existent key
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing %q, productRanks = %v\n", keyToRemove, productRanks)
}

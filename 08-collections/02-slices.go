package main

import "fmt"

func main() {
	/*
		var nos []int
		nos = append(nos, 3)
		nos = append(nos, 1)
		nos = append(nos, 4)
		nos = append(nos, 2)
		nos = append(nos, 5)
	*/

	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating a slice [using indexer]")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating a slice [using range]")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	nos2 := nos // a new reference to underlying array is create4d (not a copy)
	nos[0] = 9999
	fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])

	fmt.Printf("Before sorting, nos = %v\n", nos)
	sort(nos)
	fmt.Printf("After sorting, nos = %v\n", nos)

	fmt.Println("slicing")
	fmt.Printf("nos[2:4] = %v\n", nos[2:4])
	fmt.Printf("nos[2:] = %v\n", nos[2:])
	fmt.Printf("nos[:4] = %v\n", nos[:4])
}

func sort(list []int) { // no return results
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

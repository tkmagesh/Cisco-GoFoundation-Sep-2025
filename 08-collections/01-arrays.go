package main

import "fmt"

func main() {
	/*
		var nos [5]int
		nos[0] = 3
		nos[1] = 1
		nos[2] = 4
		nos[3] = 2
		nos[4] = 5

	*/
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating an array [using indexer]")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating an array [using range]")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	/*
		nos2 := nos // a copy of the 'nos' array is created
		nos[0] = 9999
		fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])
	*/

	var nos2 *[5]int
	nos2 = &nos
	fmt.Println((*nos2)[0], nos2[0]) // accessing the memory of a pointer to an array doesn't require dereferencing

	fmt.Printf("Before sorting, nos = %v\n", nos)
	sort(&nos)
	fmt.Printf("After sorting, nos = %v\n", nos)
}

func sort(list *[5]int) { // no return results
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

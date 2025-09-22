package main

import "fmt"

func main() {

	/*
		switch no := 22; no % 2 {
		case 0:
			fmt.Printf("%d is an even number\n", no)
		case 1:
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	// nested if-else

	/*
		if no := 21; no%2 == 0 {
			fmt.Println("even number")
		} else if no%2 == 1 {
			fmt.Println("odd number")
		} else {
			fmt.Println("Neither an even nor an odd")
		}
	*/

	// simulate nested-if in switch case
	switch no := 21; {
	case no%2 == 0:
		fmt.Println("even number")
	case no%2 == 1:
		fmt.Println("odd number")
	default:
		fmt.Println("Neither an even nor an odd")
	}

	// fall-through (opposite of 'break')
	switch no := 4; no {
	case 0:
		fmt.Println("no == 0")
		fallthrough
	case 1:
		fmt.Println("no <= 1")
		fallthrough
	case 2:
		fmt.Println("no <= 2")
		fallthrough
	case 3:
		fmt.Println("no <= 3")
		fallthrough
	case 4:
		fmt.Println("no <= 4")
		fallthrough
	case 5:
		fmt.Println("no <= 5")
		fallthrough
	case 6:
		fmt.Println("no <= 6")
		// fallthrough
	case 7:
		fmt.Println("no <= 7")

	}

	// fallthrough applied
	fmt.Println("fallthrough applied")
	switch plan := "SUPER"; plan {
	case "SUPREME":
		fmt.Println("[SUPREME] : Offline download allowed!")
		fallthrough
	case "SUPER":
		fmt.Println("[SUPER] : For a family of 3")
		fallthrough
	case "PRO":
		fmt.Println("[PRO] : Create your own playlist")
		fallthrough
	case "FREE":
		fmt.Println("[FREE] : Listen to music!")
	}
}

package main

import "fmt"

func main() {
	// var x interface{}
	var x any
	x = 100
	// x = 99.98
	// x = "Anim labore nostrud proident incididunt elit mollit."
	// x = true
	// x = struct{}{}
	fmt.Println(x)

	// x = 200
	x = "Ut in ex laboris aute in dolor id fugiat irure excepteur mollit in velit consectetur."
	// y := x * 2

	// unsafe
	// y := x.(int) * 2
	// fmt.Println(y)

	// safe
	// x = 100
	x = "Ut in ex laboris aute in dolor id fugiat irure excepteur mollit in velit consectetur."
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// using "type switch" for type assertion
	// x = 100
	// x = 99.98
	// x = "Anim labore nostrud proident incididunt elit mollit."
	// x = true
	// x = struct{}{}
	x = 5 + 5i
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case float64:
		fmt.Println("x is a float64, x * 0.01 =", val*0.01)
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case struct{}:
		fmt.Println("x is a struct{}")
	default:
		fmt.Println("x is of unknown type!")
	}

}

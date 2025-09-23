package main

import "fmt"

func main() {
	exec(f1) // f1 is invoked
	exec(f2) // f2 is invoked
	exec(f3)
	exec(func() {
		fmt.Println("anon fn inovked")
	})
}

/* Violates OCP (switch-case) */
/*
func exec(fnName string) {
	// invoke either f1 or f2 based on the argument
	switch fnName {
	case "f1":
		f1()
	case "f2":
		f2()
	default:
		fmt.Println("Invalid argument [fnName]")
	}
}
*/

func exec(fn func()) {
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

func f3() {
	fmt.Println("f3 invoked")
}

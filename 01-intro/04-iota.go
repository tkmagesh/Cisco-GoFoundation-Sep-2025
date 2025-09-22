package main

import "fmt"

func main() {
	/*
		const (
			Red   int = 0
			Green int = 1
			Blue  int = 2
		)
	*/

	/* const (
		Red   int = iota
		Green int = iota
		Blue  int = iota
	) */

	/*
		const (
			Red int = iota
			Green
			Blue
		)
	*/

	/*
		const (
			Red = iota
			Green
			Blue
		)
	*/

	/*
		const (
			Red = iota
			Green
			_
			Blue
		)
	*/

	const (
		Red   = (iota * 2) + 1 // (0 * 2) + 1 = 1
		Green                  // (1 * 2) + 1 = 3
		_                      // (2 * 2) + 1 = 5
		Blue                   // (3 * 2) + 1 = 7
	)
	fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)

	// Mimicking unix file permissions
	const (
		X = 1 << iota
		W
		R

		XW  = X | W
		WR  = W | R
		XWR = X | W | R
	)
	fmt.Printf("%b %b %b %b %b %b\n", X, W, XW, R, WR, XWR)
}

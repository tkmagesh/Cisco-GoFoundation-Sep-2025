package main

import (
	"fmt"
	"math"
)

// ver 1.0
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// ver 2.0
type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return 2 * (r.Height + r.Width)
}

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("x is neither a Circle not a Rectangle")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("x does not have Area() method")
	}
}
*/

func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

// ver 3.0
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

// ver 4.0

func PrintPerimeter(?){
	fmt.Println("Perimeter :", ?)
}

/* 
permeter formula
circle = 2 * pi * r
rectangle = 2 * (h + w)
square = 4 * side
*/

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)
	PrintPerimeter(c)

	r := Rectangle{Height: 12, Width: 14}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)
	PrintPerimeter(r)

	// PrintArea(100)

	s := Square{Side: 16}
	PrintArea(s)
	PrintPerimeter(s)
}

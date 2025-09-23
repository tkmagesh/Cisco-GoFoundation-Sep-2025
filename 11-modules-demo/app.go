package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tkmagesh/cisco-gofoundation-sep-2025/11-modules-demo/calculator"

	// "github.com/tkmagesh/cisco-gofoundation-sep-2025/11-modules-demo/calculator/utils"
	ut "github.com/tkmagesh/cisco-gofoundation-sep-2025/11-modules-demo/calculator/utils"
)

func run() {
	// fmt.Println("Application executed!")
	color.Red("Application executed!")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("Op Count :", calculator.OpCount())
	// fmt.Println("Is 21 odd ? :", utils.IsOdd(21))
	fmt.Println("Is 21 odd ? :", ut.IsOdd(21))
}

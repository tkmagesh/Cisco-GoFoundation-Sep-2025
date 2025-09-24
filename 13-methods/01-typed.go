package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float64
}

func (product Product) format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.id, product.name, product.cost)
}

func (product *Product) applyDiscount(discountPercent float64) { // no return values
	product.cost = product.cost * ((100 - discountPercent) / 100)
}

func main() {

	p := Product{id: 100, name: "Pen", cost: 10}

	// fmt.Println("Before applying discount, p => ", format(p))
	fmt.Println("Before applying discount, p => ", p.format())

	// applyDiscount(&p, 10) // apply 10% discount
	p.applyDiscount(10)

	// fmt.Println("After applying discount, p => ", format(p))
	fmt.Println("After applying discount, p => ", p.format())

}

package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float64
}

type PerishableProduct struct {
	Product
	expiry string
}

func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			id:   id,
			name: name,
			cost: cost,
		},
		expiry: expiry,
	}
}

func main() {
	milk := PerishableProduct{
		Product: Product{
			id:   100,
			name: "Milk",
			cost: 50,
		},
		expiry: "2 Days",
	}

	// fmt.Println(milk.Product.id, milk.Product.name, milk.Product.cost, milk.expiry)

	// The members of the Product are accessible through PerishableProduct

	fmt.Println(milk.id, milk.name, milk.cost, milk.expiry)

	rice := NewPerishableProduct(102, "Rice", 70, "2 years")
	fmt.Printf("%#v\n", rice)

	// use the format and applyDiscount on "rice" and print before and after

}

func format(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.id, product.name, product.cost)
}

func applyDiscount(product *Product, discountPercent float64) { // no return values
	product.cost = product.cost * ((100 - discountPercent) / 100)
}

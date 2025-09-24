package main

import "fmt"

// product = id, name, cost

type Product struct {
	id   int
	name string
	cost float64
}

func main() {
	/*
		var p Product
		p.id = 100
		p.name = "Pen"
		p.cost = 10
	*/
	/*
		var p Product = Product {
			id:   100,
			name: "pen",
			cost: 10,
		}
	*/

	// type inference
	/*
		var p = Product {
			id:   100,
			name: "pen",
			cost: 10,
		}
	*/

	// using :=
	p := Product{
		id:   100,
		name: "pen",
		cost: 10,
	}
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)

	// create a copy
	/*
		var p2 Product
		p2 = p
		p.cost = 12
		fmt.Printf("p.cost = %f, p2.cost = %f\n", p.cost, p2.cost)
	*/

	// use pointers to create references
	var p2 *Product
	p2 = &p
	p.cost = 12
	fmt.Printf("p.cost = %f, p2.cost = %f\n", p.cost, p2.cost)

	fmt.Println("Before applying discount, p => ", format(p))
	applyDiscount(&p, 10) // apply 10% discount
	fmt.Println("After applying discount, p => ", format(p))

	var x *Product = &Product{}
	fmt.Println(x)

	var y *Product = new(Product)
	fmt.Println(y)

	fmt.Println(new(int))
}

func format(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.id, product.name, product.cost)
}

func applyDiscount(product *Product, discountPercent float64) { // no return values
	product.cost = product.cost * ((100 - discountPercent) / 100)
}

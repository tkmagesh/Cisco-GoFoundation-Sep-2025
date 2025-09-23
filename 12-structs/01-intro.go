package main

import "fmt"

// product = id, name, cost

func main() {
	/*
		var p struct {
			id   int
			name string
			cost float64
		}
		p.id = 100
		p.name = "Pen"
		p.cost = 10
	*/
	/*
		var p struct {
			id   int
			name string
			cost float64
		} = struct {
			id   int
			name string
			cost float64
		}{
			id:   100,
			name: "pen",
			cost: 10,
		}
	*/

	// type inference
	/*
		var p = struct {
			id   int
			name string
			cost float64
		}{
			id:   100,
			name: "pen",
			cost: 10,
		}
	*/

	// using :=
	p := struct {
		id   int
		name string
		cost float64
	}{
		id:   100,
		name: "pen",
		cost: 10,
	}
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)

	// create a copy
	/*
		var p2 struct{ id  int; name  string; cost  float64;}
		p2 = p
		p.cost = 12
		fmt.Printf("p.cost = %f, p2.cost = %f\n", p.cost, p2.cost)
	*/

	// use pointers to create references
	var p2 *struct {
		id   int
		name string
		cost float64
	}
	p2 = &p
	p.cost = 12
	fmt.Printf("p.cost = %f, p2.cost = %f\n", p.cost, p2.cost)

	fmt.Println("Before applying discount, p = ", format( /* ? */ ))
	applyDiscount( /* ? */ ) // apply 10% discount
	fmt.Println("After applying discount, p = ", format( /* ? */ ))
}

func format( /* product */ ) string {
	// Id = 100, Name = "pen", Cost = 10
	// return ?
}

func applyDiscount( /* ? */ ) { // no return values
	// update the give product cost after applying the given discount
}

package main

import (
	"fmt"
	"strings"
)

type Product struct {
	id   int
	name string
	cost float64
}

func NewProduct(id int, name string, cost float64) *Product {
	return &Product{
		id:   id,
		name: name,
		cost: cost,
	}
}

func (product Product) format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.id, product.name, product.cost)
}

func (product *Product) applyDiscount(discountPercent float64) { // no return values
	product.cost = product.cost * ((100 - discountPercent) / 100)
}

/*
Create a type for "ShoppingCart"
- should be able to add products to the cart and number of units for each product
- apply discount on the over all cart value
- print the cart (list the products and display the cart value)
*/

type ShoppingCart struct {
	lineItems []*LineItem
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (cart *ShoppingCart) addItem(p *Product, units int) {
	cart.lineItems = append(cart.lineItems, NewLineItem(p, units))
}

func (cart *ShoppingCart) format() string {
	var builder strings.Builder
	for _, li := range cart.lineItems {
		builder.WriteString(fmt.Sprintf("%s\n", li.format()))
	}
	builder.WriteString(fmt.Sprintf("Cart Value : %0.2f\n", cart.getValue()))
	return builder.String()
}

func (cart ShoppingCart) getValue() float64 {
	var cartValue float64
	for _, li := range cart.lineItems {
		cartValue += li.getValue()
	}
	return cartValue
}

type LineItem struct {
	product *Product
	units   int
}

func NewLineItem(p *Product, units int) *LineItem {
	return &LineItem{
		product: p,
		units:   units,
	}
}

func (li LineItem) format() string {
	return fmt.Sprintf("%s, Units = %d", li.product.format(), li.units)
}

func (li LineItem) getValue() float64 {
	return li.product.cost * float64(li.units)
}

func main() {
	cart := NewShoppingCart()
	cart.addItem(NewProduct(100, "Pen", 10), 20)
	cart.addItem(NewProduct(101, "Pencil", 5), 40)
	cart.addItem(NewProduct(103, "Marker", 50), 5)
	fmt.Println(cart.format())
}

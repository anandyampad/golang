package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

/*
	Write the following functions for "Product"
	Format() will return a formatted string of the given "product"
	ApplyDiscount() will apply the given discount % on the given "product"
*/

/*
	Use the above functions for a PerishableProduct (grapes)
*/

func FormatProduct(product Product) string {
	return fmt.Sprintf("[Id:%d Name:%s Cost:%0.2f]", product.Id, product.Name, product.Cost)
}

func ApplyDiscount(product *Product, discount float32) {
	product.Cost -= (product.Cost * discount) / 100
}

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	Product
	// Dummy
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	// grapes := PerishableProduct{Product{100, "Grapes", 50},  "2 Days"}
	/*
		grapes := PerishableProduct{
			Product: Product{100, "Grapes", 50},
			Expiry:  "2 Days",
		}
	*/
	grapes := NewPerishableProduct(100, "Grapes", 50, "2 Days")
	fmt.Println(grapes.Product.Name)
	// fmt.Println(grapes.Name)

	fmt.Println(FormatProduct(grapes.Product))
	ApplyDiscount(&grapes.Product, 20)
	fmt.Println(grapes)
}

package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s",
		p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

func (products Products) String() string {
	sb := strings.Builder{}
	for _, p := range products {
		sb.WriteString(fmt.Sprintf("%s\n", p.String()))
	}
	return sb.String()
}

/*
Write the apis for the following
*/

/*
IndexOf => return the index of the given product

	ex:  returns the index of the given product
*/
func (products Products) IndexOf(product Product) (int, error) {
	for i, val := range products {
		if val == product {
			return i, nil
		}
	}
	err := errors.New("Product not found!")
	return 0, err
}

/*
Includes => return true if the given product is present in the collection else return false
	ex:  returns true if the given product is present in the collection
*/

func (products Products) Includes(product Product) bool {
	for _, val := range products {
		if product == val {
			return true
		}
	}
	return false
}

/*
Filter => return a new collection of products that satisfy the given condition
	use cases:
		1. filter all costly products (cost > 1000)
			OR
		2. filter all stationary products (category = "Stationary")
			OR
		3. filter all costly stationary products
		etc
*/

type Predicate func(product Product) bool

func (products Products) Filter(predicate Predicate) Products {
	result := make(Products, 0)
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

/*
	Any => return true if any of the product in the collections satifies the given criteria
		use cases:
			1. are there any costly product (cost > 1000)?
			2. are there any stationary product (category = "Stationary")?
			3. are there any costly stationary product?
			etc
*/

func (products Products) Any(predicate Predicate) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

/*
	All => return true if all the products in the collections satifies the given criteria
		use cases:
			1. are all the products costly products (cost > 1000)?
			2. are all the products stationary products (category = "Stationary")?
			3. are all the products costly stationary products?
			etc
*/

func (products Products) All(predicate Predicate) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

/*

	Write the apis for the following
		Sort => Sort the products collection by any attribute
			IMPORTANT : MUST Use sort.Sort() function
            use cases:
                1. sort the products collection by cost
                2. sort the products collection by name
                3. sort the products collection by units
				etc

*/

// By Id
func (p Products) Len() int           { return len(p) }
func (p Products) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Products) Less(i, j int) bool { return p[i].Id < p[j].Id }

// By Cost
type ByCost struct {
	Products
}

func (p ByCost) Less(i, j int) bool { return p.Products[i].Cost < p.Products[j].Cost }

// By Name
type ByName struct {
	Products
}

func (p ByName) Less(i, j int) bool { return p.Products[i].Name < p.Products[j].Name }

// By Units
type ByUnits struct {
	Products
}

func (p ByUnits) Less(i, j int) bool { return p.Products[i].Units < p.Products[j].Units }

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	//IndexOf
	PencilProduct := Product{107, "Pencil", 2, 100, "Stationary"}
	index, err := products.IndexOf(PencilProduct)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Index of Pencil Product=", index)
	}

	//Includes
	StoveProduct := Product{102, "Stove", 5000, 5, "Utencil"}
	includes := products.Includes(StoveProduct)
	if includes {
		fmt.Println("Stove Product is included")
	} else {
		fmt.Println("Stove Product is NOT included")
	}

	//Filter
	fmt.Println("\nProducts with Cost > 1000:")
	costlyProductPredicate := func(product Product) bool {
		if product.Cost > 1000 {
			return true
		} else {
			return false
		}
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println(costlyProducts)

	fmt.Println("\nProducts with Stationary Category:")
	StationaryProductsPredicate := func(product Product) bool {
		if product.Category == "Stationary" {
			return true
		} else {
			return false
		}
	}

	StationaryProducts := products.Filter(StationaryProductsPredicate)
	fmt.Println(StationaryProducts)

	fmt.Println("\nProducts with Cost > 1000 and having Stationary Category:")
	CostlyStationaryProductsPredicate := func(product Product) bool {
		if StationaryProductsPredicate(product) && costlyProductPredicate(product) {
			return true
		} else {
			return false
		}
	}

	CostlyStationaryProducts := products.Filter(CostlyStationaryProductsPredicate)
	fmt.Println(CostlyStationaryProducts)

	//Any
	anyCostlyProducts := products.Any(costlyProductPredicate)
	fmt.Print("\nAre there any costly(>1000) Products:")
	fmt.Println(anyCostlyProducts)

	anyStationaryProducts := products.Any(StationaryProductsPredicate)
	fmt.Print("\nAre there any Stationary Products:")
	fmt.Println(anyStationaryProducts)

	anyCostlyStationaryProducts := products.Any(CostlyStationaryProductsPredicate)
	fmt.Print("\nAre there any Costly Stationary Products:")
	fmt.Println(anyCostlyStationaryProducts)

	// All
	allCostlyProducts := products.All(costlyProductPredicate)
	fmt.Print("\nAre there all costly(>1000) Products:")
	fmt.Println(allCostlyProducts)

	allStationaryProducts := products.All(StationaryProductsPredicate)
	fmt.Print("\nAre there all Stationary Products:")
	fmt.Println(allStationaryProducts)

	allCostlyStationaryProducts := products.All(CostlyStationaryProductsPredicate)
	fmt.Print("\nAre there all Costly Stationary Products:")
	fmt.Println(allCostlyStationaryProducts)

	//Sort
	sort.Sort(ByCost{products})
	fmt.Println("\nAfter sorting by Cost")
	fmt.Println(products)

	sort.Sort(ByName{products})
	fmt.Println("\nAfter sorting by Name")
	fmt.Println(products)

	sort.Sort(ByUnits{products})
	fmt.Println("\nAfter sorting by Units")
	fmt.Println(products)
}

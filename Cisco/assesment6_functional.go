package main

import (
	"errors"
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

/*
Write the apis for the following
*/

/*
IndexOf => return the index of the given product

	ex:  returns the index of the given product
*/
func (products Products)IndexOf(product Product) (int, error) {
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

/*
type Specification interface {
	IsSatisfied(product *Product) bool
}

type CostSpecification struct {
	Cost float32
}

func (c CostSpecification) IsSatisfied(product *Product) bool {
	if product.Cost > c.Cost {
		return true
	} else {
		return false
	}
}

type CategorySpecification struct {
	Category string
}

func (c CategorySpecification) IsSatisfied(product *Product) bool {
	if product.Category == c.Category {
		return true
	} else {
		return false
	}
}

type CostAndCategorySpecification struct {
	Cost     float32
	Category string
}

func (cc CostAndCategorySpecification) IsSatisfied(product *Product) bool {
	if (product.Cost > cc.Cost) && (product.Category == cc.Category) {
		return true
	} else {
		return false
	}
}
*/

type Predicate func(product Product)bool
func Filter(products Products)(predicate Predicate) []Product {
	result := make([]Product, 0)
	for _, value := range products {
		if predicate(&value) {
			result = append(result, value)
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

func (products Products)Any(predicate Predicate) bool {
	for _, v := range products {
		if predicate(&v) {
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

func (products Products)All(predicate Predicate) bool {
	for _, v := range products {
		if !Predicate(&v) {
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


type SortProducts interface{
	
}


func (p Products) Len() int           { return len(p) }
func (p Products) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Products) Less(i, j int) bool { return p[i].Id< p[j].Id }


type ByCost struct{
	Products
}

func (p ByCost) Less(i, j int) bool { return p[i].Cost< p[j].Cost }

func main() {
	products := []Product{
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
	index, err := IndexOf(products, PencilProduct)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Index of Pencil Product=", index)
	}

	//Includes
	StoveProduct := Product{102, "Stove", 5000, 5, "Utencil"}
	includes := Includes(products, StoveProduct)
	if includes {
		fmt.Println("Stove Product is included")
	} else {
		fmt.Println("Stove Product is NOT included")
	}

	//Filter
	costSpec := CostSpecification{1000}
	costlyProducts := Filter(products, costSpec)

	fmt.Println("\nProducts with Cost > 1000:")
	fmt.Println(costlyProducts)

	categorySpec := CategorySpecification{"Stationary"}
	StationaryProducts := Filter(products, categorySpec)

	fmt.Println("\nProducts with Stationary Category:")
	fmt.Println(StationaryProducts)

	costAndCategorySpec := CostAndCategorySpecification{1000, "Stationary"}
	CostlyStationaryProducts := Filter(products, costAndCategorySpec)

	fmt.Println("\nProducts with Cost > 1000 and having Stationary Category:")
	fmt.Println(CostlyStationaryProducts)

	//Any
	anyCostlyProducts := Any(products, costSpec)
	fmt.Print("\nAre there any costly(>1000) Products:")
	fmt.Println(anyCostlyProducts)

	anyStationaryProducts := Any(products, categorySpec)
	fmt.Print("\nAre there any Stationary Products:")
	fmt.Println(anyStationaryProducts)

	anyCostlyStationaryProducts := Any(products, costAndCategorySpec)
	fmt.Print("\nAre there any Costly Stationary Products:")
	fmt.Println(anyCostlyStationaryProducts)

	// All
	allCostlyProducts := All(products, costSpec)
	fmt.Print("\nAre there all costly(>1000) Products:")
	fmt.Println(allCostlyProducts)

	allStationaryProducts := All(products, categorySpec)
	fmt.Print("\nAre there all Stationary Products:")
	fmt.Println(allStationaryProducts)

	allCostlyStationaryProducts := All(products, costAndCategorySpec)
	fmt.Print("\nAre there all Costly Stationary Products:")
	fmt.Println(allCostlyStationaryProducts)

	//Sort
	sort.Sort(ByCost(products))
	fmt.Println("After sorting by Cost")

	fmt.Println(products)

	sort.Sort(ByName(products))
	fmt.Println("After sorting by Name")

	fmt.Println(products)

	sort.Sort(ByUnits(products))
	fmt.Println("After sorting by Units")

	fmt.Println(products)

}

package main

import "fmt"

type Employee struct {
	Name string
}

func main() {
	//var x interface{}
	var x any
	x = 100
	x = "hello world"
	x = false
	//x = 10.00
	//x = struct{}{}
	//x = Employee{"Anand"}
	//x = 500

	//x = struct{ Id int }{Id: 100}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 300 =", val+300)
	case float64:
		fmt.Println("x is a float, 90% of x is =", val*0.9)
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case struct{}:
		fmt.Println("x is a struct{}")
	case Employee:
		fmt.Println("x is an Employee, Name = ", val.Name)
	default:
		fmt.Println("Unknown type")
	}
}

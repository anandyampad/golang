package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {

	var emp Employee
	PrintEmployee(emp)

	emp1 := Employee{
		Id:     100,
		Name:   "Anand",
		Salary: 10000,
	}
	PrintEmployee(emp1)

	emp2 := emp1
	emp2.Salary = 30000

	PrintEmployee(emp1)
	PrintEmployee(emp2)
}

func PrintEmployee(e Employee) {
	fmt.Println(e)
}

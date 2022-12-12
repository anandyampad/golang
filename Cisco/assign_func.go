package main

import "fmt"

func main() {

	var greet1 func()
	greet1 = func() { //greet1 := func() {
		fmt.Println("Hi there!")
	}

	var greet2 func(string)
	greet2 = func(userName string) { //greet2 := func(userName string) {
		fmt.Printf("Hi %q, Have a nice day!\n", userName)
	}

	var greet3 func(string) string
	greet3 = func(userName string) string { //greet3 := func(userName string) string {
		var msg string
		msg = fmt.Sprintf("Hi %q, Have a nice day", userName)
		return msg
	}

	var result1 func(int, int) int
	result1 = func(x, y int) int { //result1 := func(x, y int) int {
		return x + y
	}

	var result2 func(int, int) (int, int)
	result2 = func(x, y int) (quotient, remainder int) { //result2 := func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}

	greet1()
	greet2("Anand")
	fmt.Println(greet3("Anand"))
	fmt.Println(result1(100, 200))
	q, r := result2(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}

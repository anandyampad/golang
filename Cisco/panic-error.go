package main

import (
	"fmt"
	"runtime/debug"
)

func main() {

	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("Somethign went wrong")
			fmt.Println(e)
			debug.PrintStack()
		}
	}()

	if divideClient(100, 0) == 0 {
		fmt.Println("Error occured")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("test")
	}

}

func divideClient(x, y int) int {
	z := 10
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println(e)
			z = 0
		}
	}()

	fmt.Println("result=", divide(x, y))
	return z
}
func divide(x, y int) int {
	return x / y
	//fmt.Println("Result=", x/y)
}

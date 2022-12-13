package main

import "fmt"

func main() {

	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	x := 5
	//we can only defer functions
	fmt.Println(x)
	defer fmt.Println()
	fmt.Println("f1 started")
	{
		fmt.Println("Function started")
		defer func() {
			fmt.Println("f1 deffered")
		}()
		fmt.Println("Function completed")
	}

	fmt.Println("f1 completed")
}

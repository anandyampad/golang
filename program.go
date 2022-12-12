package main

import "fmt"

func main() {
	print("Hello World!\n")

	var x, y int = 10, 20
	sum := x + y

	fmt.Printf("sum of %d and %d is %d\n", x, y, sum)

	const (
		red = iota + 4
		blue
		green
	)
	fmt.Printf("red = %d blue = %d green = %d\n", red, blue, green)

}

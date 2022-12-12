package main

import "fmt"

func main() {

	for {
		userchoice := ReadUserChoice()

		switch userchoice {
		case 1:
			x, y := ReadOperands()
			fmt.Printf("Result = %d\n", Add(x, y))
		case 2:
			x, y := ReadOperands()
			fmt.Printf("Result = %d\n", Subtract(x, y))
		case 3:
			x, y := ReadOperands()
			fmt.Printf("Result = %d\n", Multiply(x, y))
		case 4:
			x, y := ReadOperands()
			q, r := Divide(x, y)
			fmt.Printf("Quotient = %d Remainder = %d \n", q, r)
		case 5:
			fmt.Println("Thank you")
		default:
			fmt.Println("Invalid Choice, Try again")
		}

		if userchoice == 5 {
			break
		}
	}
}

func ReadUserChoice() int {
	var choice int
	fmt.Println("1. Add\n2. Subtract\n3. Multiply\n4. Divide\n5. Exit")
	fmt.Print("Enter your chocie:")
	fmt.Scanf("%d", &choice)
	fmt.Println()
	return choice
}

func ReadOperands() (int, int) {
	var x, y int
	fmt.Println("Enter two numbers seperated by space")
	fmt.Scanf("%d%d", &x, &y)
	return x, y
}

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Multiply(x, y int) int {
	return x * y
}

func Divide(x, y int) (int, int) {
	return x / y, x % y
}

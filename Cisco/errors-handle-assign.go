package main

import (
	"errors"
	"fmt"
)

func main() {

	for {
		userchoice, err := ReadUserChoice()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if userchoice == 5 {
			fmt.Println("Thank you")
			break
		}
		err = DoOperations(userchoice)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func DoOperations(userchoice int) error {
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
		q, r, err := Divide(x, y)
		if err != nil {
			return err
		}
		fmt.Printf("Quotient = %d Remainder = %d \n", q, r)
	}
	return nil
}

func ReadUserChoice() (choice int, err error) {
	fmt.Println("1. Add\n2. Subtract\n3. Multiply\n4. Divide\n5. Exit")
	fmt.Print("Enter your chocie:")
	fmt.Scan(&choice)
	fmt.Println()
	if choice < 1 || choice > 5 {
		err = errors.New("Inavlid choice")
		return
	}
	return
}

func ReadOperands() (int, int) {
	var x, y int
	fmt.Println("Enter two numbers seperated by space")
	fmt.Scan(&x, &y)
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

func Divide(x, y int) (q int, r int, err error) {
	if y == 0 {
		err = errors.New("Divide by Zero error!")
		return
	}
	q = x / y
	r = x % y
	return
}

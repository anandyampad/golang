package main

import "fmt"

func main() {

	var choice int

	for {

		fmt.Println("1. Add\n2. Subtract\n3. Multiply\n4. Divide\n5. Exit")
		fmt.Print("Enter your chocie:")
		fmt.Scanf("%d", &choice)
		fmt.Println()

		var x, y int

		switch choice {

		case 1:
			fmt.Println("Enter two numbers seperated by space")
			fmt.Scanf("%d%d", &x, &y)
			fmt.Printf("Result = %d\n", x+y)
		case 2:
			fmt.Println("Enter two numbers seperated by space")
			fmt.Scanf("%d%d", &x, &y)
			fmt.Printf("Result = %d\n", x-y)
		case 3:
			fmt.Println("Enter two numbers seperated by space")
			fmt.Scanf("%d%d", &x, &y)
			fmt.Printf("Result = %d\n", x*y)
		case 4:
			fmt.Println("Enter two numbers seperated by space")
			fmt.Scanf("%d%d", &x, &y)
			fmt.Printf("Result = %d\n", x/y)
		case 5:
			fmt.Println("Thank you")
		default:
			fmt.Println("Invalid Choice, Try again")
		}

		if choice == 5 {
			break
		}
	}

}

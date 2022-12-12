package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number is ", i)
	}

	for sum := 1; sum < 100; {
		fmt.Printf("%d is less than 100\n", sum)
		sum += sum
	}

	sum := 1
	for {
		fmt.Printf("%d is less than 100\n", sum)
		sum += sum
		if sum > 100 {
			break
		}
	}

LOOP:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("i=%d j=%d\n", i, j)
			if i == j {
				fmt.Println("==================")
				continue LOOP
			}
		}
	}
}

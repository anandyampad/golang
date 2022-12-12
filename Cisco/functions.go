package main

import "fmt"

func main() {

	q, _ := divide(100, 7)
	fmt.Printf("100 divide by 7 is %d\n", q)
	fmt.Printf("Sum = %d\n", sum(10, 20))
	fmt.Printf("Sum = %d\n", sum())
	fmt.Printf("Sum = %d\n", sum(1, 4, 2, 6, 8))
}

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return quotient, remainder
}

func sum(nums ...int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}
	return ans
}

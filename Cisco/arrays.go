package main

import "fmt"

func main() {
	myArr := [4]int{1, 2, 3, 4}
	fmt.Println(myArr, len(myArr))

	for i, val := range myArr {
		val = val * 10
		fmt.Printf("%d -> %d\n", i, val)
	}

	fmt.Println(myArr)
	for i := 0; i < len(myArr); i++ {
		myArr[i] = myArr[i] * 10
	}

	fmt.Println(myArr)
}

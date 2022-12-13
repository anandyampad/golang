package main

import "fmt"

func main() {
	nos := []int{1, 2}
	fmt.Println(nos)

	newNums := nos
	fmt.Println(nos, newNums)
	newNums = append(newNums, 6)
	//nos[5] wont work
	fmt.Println(nos, newNums)

	fmt.Println("cap nos=", cap(nos), "cap newNums=", cap(newNums))
	fmt.Println(cap(nos), cap(newNums))

	var mySlice []int = make([]int, 5)
	fmt.Println(mySlice)

	var mySlice1 []int = make([]int, 5, 10)
	fmt.Println("len(mySlice1) =", len(mySlice1), "cap(mySlice1) =", cap(mySlice1))
}

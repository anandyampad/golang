package main

import "fmt"

func main() {

	ch := make(chan int)
	go add(100, 200, ch)
	res := <-ch
	fmt.Printf("result=%d\n", res)
}

func add(x, y int, ch chan int) {
	ans := x + y
	ch <- ans
	for i := 0; i < 500; i++ {
		fmt.Print(i) /////
	}
}

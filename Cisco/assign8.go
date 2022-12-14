package main

import "fmt"

/*
Write a goroutine to generate N fibonocci numbers (N is the input)
Print the generated fibonocci numbers in the main function
*/

func main() {

	ch := generateNums(10)

	for {

		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			continue
		}
		break
	}
}

func generateNums(n int) chan int {
	ch := make(chan int)
	go func() {

		n1 := 0
		n2 := 1

		for i := 0; i < n; i++ {
			ch <- n1
			n1, n2 = n2, n1+n2
		}
		close(ch)

	}()
	return ch
}

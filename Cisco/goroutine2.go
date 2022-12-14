package main

import (
	"fmt"
)

func main() {

	ch := generateNums()
	for {
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			continue
		}
		break
	}
}

func generateNums() chan int {

}

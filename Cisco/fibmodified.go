/*
Write a goroutine to generate N fibonocci numbers (N is the input)
Print the generated fibonocci numbers in the main function
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generateFibonocci()
	for no := range ch {
		fmt.Println(no)
	}

}

func generateFibonocci() <-chan int {
	ch := make(chan int)
	ch1 := make(chan struct{})

	go func() {
		fmt.Scanln()
		ch1 <- struct{}{}
		close(ch1)
	}()
	go func() {
		x, y := 0, 1
	Loop:
		for {
			select {
			case <-ch1:
				break Loop
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)

	}()
	return ch
}

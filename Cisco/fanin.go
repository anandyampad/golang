package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	oddCh := generateOddNums()
	evenCh := generateEvenNums()

	resultCh := fanIn(oddCh, evenCh)
	for data := range resultCh {
		fmt.Println(data)
	}

}

func fanIn(chs ...chan int) chan int {
	resultChan := make(chan int)

	func() {
		wg := sync.WaitGroup{}
		for _, ch := range chs {
			wg.Add(1)
			func(ch1 chan int) {
				for data := range ch1 {
					resultChan <- data
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}

func generateOddNums() chan int {
	ch := make(chan int)
	func() {
		for i := 1; i < 10; i++ {
			time.Sleep(300 * time.Millisecond)
			ch <- (2 * i) - 1
		}
		close(ch)
	}()
	return ch
}

func generateEvenNums() chan int {
	ch := make(chan int)
	func() {
		for i := 1; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- 2 * i
		}
		close(ch)
	}()
	return ch
}

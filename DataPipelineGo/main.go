package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var readFile int32 = 0

func main() {

	ch := make(chan int)          //read from file and write to ch1
	oddChan := make(chan int)     // send odd data here
	evenChan := make(chan int)    // send even data here
	oddSumChan := make(chan int)  //send odd sum here
	evenSumChan := make(chan int) //send even sum here

	wg.Add(6)
	go readDataFile("data1.dat", ch)
	go readDataFile("data2.dat", ch)
	go splitFileData(ch, oddChan, evenChan)
	go processFileData(oddChan, oddSumChan)
	go processFileData(evenChan, evenSumChan)
	go mergeFileData(oddSumChan, evenSumChan)

	wg.Wait()
	fmt.Println("Main completed")
}

// Open file
// Read from  file
// send it to ch channel
func readDataFile(fileName string, ch chan int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, er := strconv.Atoi(scanner.Text())
		if er != nil {
			log.Println("can not convert to int")
		} else {
			ch <- val
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	atomic.AddInt32(&readFile, 1)
	wg.Done()
}

// Receive data from ch from both senders
// Process the data to differentate odd or even
// Send data to seperate oddChan and evenChan channels
func splitFileData(ch, chOdd, chEven chan int) {
	for {
		time.Sleep(5 * time.Millisecond) //without this it goes into deadlock
		if val := atomic.LoadInt32(&readFile); val != 2 {
			data := <-ch
			if data%2 == 0 {
				chEven <- data
			} else {
				chOdd <- data
			}
			continue
		} else {
			break
		}
	}
	close(ch)
	close(chOdd)
	close(chEven)
	wg.Done()
}

//Receive data from oddChan/evenChan
//Calculate sum of oddChan/evenChan Channel data
//Send data to oddSumChan/evenSumChan

func processFileData(ch, sumch chan int) {
	sum := 0
	for data := range ch {
		sum += data
	}
	sumch <- sum
	wg.Done()
}

// Receive data in oddSumChan/evenSumChan
// write ii to file
func mergeFileData(oddChSum, evenChSum chan int) {
	var localwg sync.WaitGroup
	oddSum := 0
	evenSum := 0
	localwg.Add(2)
	go func() {
		defer localwg.Done()
		oddSum = <-oddChSum
		close(oddChSum)

	}()
	go func() {
		defer localwg.Done()
		evenSum = <-evenChSum
		close(evenChSum)
	}()
	localwg.Wait()
	fmt.Println("oddsum=", oddSum, " evensum=", evenSum)
	WriteToFile(oddSum, evenSum)
	wg.Done()
}

func WriteToFile(oddsum, evensum int) {
	f, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	len, err := f.WriteString(fmt.Sprintf("Even Total : %d\nOdd Total : %d\n", oddsum, evensum))
	if err != nil {
		log.Fatal(len, err)
	}
	f.Sync()
}

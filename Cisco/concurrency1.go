package main

import "fmt"

func main() {
	fmt.Println("main started")
	go f1()
	go f2()
	go f3()
	go f4()
	go f5()
	for i := 0; i < 100; i++ {
		fmt.Print(i)
	}
	fmt.Println("main completed")

}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
func f3() {
	fmt.Println("f3 invoked")
}
func f4() {
	fmt.Println("f4 invoked")
}
func f5() {
	fmt.Println("f5 invoked")
}

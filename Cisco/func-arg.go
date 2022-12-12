package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
)

type ArithOperation func(int, int)

func main() {

	//logadd := logOperation(add)
	//logadd(30, 20)
	logOperation(add)(10, 20)
	logOperation(subtract)(20, 10)
	logOperation(multiply)(20, 10)
	logOperation(divide)(20, 10)

	fmt.Println("fn1=", increment()())
	fmt.Println("fn2=", increment()())
	fmt.Println("fn3=", increment()())

}

func increment() func() int {
	i := 0
	fn := func() int {
		i++
		return i
	}
	return fn
}

func logOperation(fn ArithOperation) ArithOperation {

	functionName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	return func(x, y int) {
		log.Println("[", functionName, "]:Operation started")
		fn(x, y)
		log.Println("[", functionName, "]:Operation completed")
		log.Println("------------------------------------------")
	}
}

func add(x, y int) {
	log.Println("Result=", x+y)
}

func subtract(x, y int) {
	log.Println("Result=", x-y)
}
func multiply(x, y int) {
	log.Println("Result=", x*y)
}
func divide(x, y int) {
	log.Println("Result=", x/y)
}

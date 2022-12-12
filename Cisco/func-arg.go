package main

import (
	"log"
	"reflect"
	"runtime"
)

func main() {
	logOperation(20, 10, add)
	logOperation(20, 10, subtract)
	logOperation(20, 10, multiply)
	logOperation(20, 10, divide)
}

func logOperation(x, y int, fn func(int, int) int) {
	functionName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	log.Println("[", functionName, "]:Operation started")
	log.Println("Result=", fn(x, y))
	log.Println("[", functionName, "]:Operation completed")
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
func multiply(x, y int) int {
	return x * y
}
func divide(x, y int) int {
	return x / y
}

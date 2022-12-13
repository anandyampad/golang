package main

import "fmt"

/*prime numbers
write a function "genPrimes" that generates the prime numbers from the given start to end and returns the primeNos
the main function receives the prime nos and prints them
*/

func main() {

	primes := genPrime(2, 11)
	fmt.Println(primes)
}

func genPrime(start, end int) []int {
	var primes []int
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}

	}
	return primes
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10, 20))                                 //=> 30
	fmt.Println(sum(10, 20, 30, 40))                         //=> 100
	fmt.Println(sum(10))                                     //=> 10
	fmt.Println(sum())                                       //=> 0
	fmt.Println(sum(10, "20", 30, "40"))                     //=> 100 (use strconv.Atoi())
	fmt.Println(sum(10, "20", 30, "40", "abc"))              //=> 100
	fmt.Println(sum(10, 20, []int{30, 40}))                  //=> 100
	fmt.Println(sum(10, 20, []any{30, 40, []int{10, 20}}))   //=> 130
	fmt.Println(sum(10, 20, []any{30, 40, []any{10, "20"}})) //=> 130
}

func sum(nums ...any) int {
	ans := 0
	for _, x := range nums {

		switch val := x.(type) {
		case int:
			ans += val
		case []int:
			for _, num := range val {
				ans += num
			}
		case string:
			if n, err := strconv.Atoi(val); err == nil {
				ans += n
			}
		case []any:
			ans += sum(val...)
		}
	}
	return ans
}

package main

import (
	"fmt"
	"strings"
)

func main() {

	//find the most occuring "word size" and print the "word size" and the number of occurences
	//ex output: 4 letter words occurs the most with 30 occurrences
	//hint : use strings.Split() function

	str := "Pariatur nisi nostrud id ipsum esse minim velit Lorem eiusmod est reprehenderit sint Esse non proident labore dolore labore nisi dolore dolor proident voluptate minim laborum Ad irure ut Lorem exercitation Ipsum reprehenderit consequat veniam non reprehenderit ut eiusmod magna magna aliquip ut sint Mollit irure fugiat nostrud non excepteur aliquip cillum Labore dolore deserunt irure ea tempor Consequat ea quis ipsum esse minim reprehenderit amet Qui do dolor do anim occaecat culpa commodo sunt mollit excepteur laboris tempor Ullamco labore quis culpa laborum sunt Culpa incididunt ad consectetur ut deserunt tempor proident ut Qui quis aliqua fugiat sunt dolore commodo reprehenderit veniam tempor Voluptate sit laboris sunt et do sint ea irure veniam duis laborum Quis irure id officia inNisi velit proident cupidatat laborum velit enim qui deserunt consectetur Ad ea quis veniam cupidatat tempor sunt elit velit Sit anim irure sunt ut exercitation excepteur elit laboris occaecat dolor Duis ad commodo ut quis nisi pariatur anim Fugiat ad id anim velit labore"

	values := strings.Split(str, " ")

	count := 0
	myAns := ""
	mymap := make(map[string]int)
	for _, val := range values {
		if _, exist := mymap[val]; exist {
			mymap[val] = mymap[val] + 1
		} else {
			mymap[val] = 1
		}
		if mymap[val] > count {
			myAns = val
			count = mymap[val]
		}
	}

	fmt.Printf("%d letter word occured %d times\n", len(myAns), count)
}

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)

	var keyword string
	if num % 2 == 0 {
		keyword = "even"
	} else {
		keyword = "odd"
	}

	fmt.Printf("%d is %s", num, keyword)
}


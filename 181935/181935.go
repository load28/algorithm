package main

import (
	"fmt"
)

func main() {
	result := solution(10)
	fmt.Println(result)
}

func solution(n int) int {
	if n%2 == 0 {
		return calculateProduct(n)
	} else {
		return calculateSum(n)
	}
}

func calculateProduct(n int) int {
	product := n * (n + 1) * (2*n + 1) / 6
	return product
}

func calculateSum(n int) int {
	sum := n * (n + 1) / 2
	return sum
}

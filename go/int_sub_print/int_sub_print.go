package main

import "fmt"

func main() {
	var a, b, sum int

	fmt.Scan(&a, &b)
	sum = a + b
	fmt.Printf("%d + %d = %d", a, b, sum)
}
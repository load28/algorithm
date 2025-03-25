package main

import "fmt"

func main() {
	var str string
	var num int
	var sum_str string

	fmt.Scan(&str, &num)

	for i := 0; i < num; i ++ {
		sum_str += str
	}

	fmt.Println(sum_str)
}
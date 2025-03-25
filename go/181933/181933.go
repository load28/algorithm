package main

import "fmt"

func main() {
	result := solution(1, 11, true)
	fmt.Println(result)
}

func solution(a int, b int, flag bool) int {
	if flag {
		return a + b
	} else {
		return a - b
	}
}

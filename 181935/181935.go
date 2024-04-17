package main

import "fmt"

func main() {
	var result = solution(10)
	fmt.Println(result)
}

// https://school.programmers.co.kr/learn/courses/30/lessons/181935
func solution(n int) int {
	return some(n)
}

func some(n int) int {
	if n <= 1 {
		return n
	}

	isEven := n%2 == 0
	if isEven {
		return n * some(n-2)
	} else {
		return n + some(n-2)
	}
}

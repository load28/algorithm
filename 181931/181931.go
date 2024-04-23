package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/181931
func main() {
	arr := []bool{true, true, false}
	result := solution(1, 3, arr)
	fmt.Println(result)
}

func solution(a int, d int, included []bool) int {
	result := 0
	for index, value := range included {
		if value {
			result += a + (d * index)
		}
	}
	return result
}

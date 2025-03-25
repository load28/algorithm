package main 

import (
	"fmt"
	"strconv"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/181939
func main() {
	r := solution(10, 11)
	fmt.Println(r)
}

func solution(a int, b int) int {
	result1, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	result2, _ := strconv.Atoi(strconv.Itoa(b) + strconv.Itoa(a))

	if result1 == result2 {
		return result1
	} else if result1 < result2 {
		return result2
	} else {
		return result1
	}
}

package main

import (
	"fmt" 
	"strconv"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/181938
func main() {
	r := solution(10, 1)
	fmt.Println(r)
}

func solution(a int, b int) int {
	sumNum, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	timeNum := 2 * a * b

	if sumNum == timeNum || sumNum > timeNum {
		return sumNum
	} else {
		return timeNum
	}
}
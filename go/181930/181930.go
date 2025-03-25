// https://school.programmers.co.kr/learn/courses/30/lessons/181930

package main

import (
	"fmt"
	"math"
)

func main() {
	result := solution(1, 1, 2)
	fmt.Println(result)
}

func solution(a int, b int, c int) int {
	if a != b && b != c && c != a {
		fmt.Println("1")
		return sum(1, a, b, c)
	} else if a == b && b == c && c == a {
		fmt.Println("3")
		return sum(3, a, b, c)
	} else {
		fmt.Println("2")
		return sum(2, a, b, c)
	}
}

func sum(sameCount int, a int, b int, c int) int {
	result := 1.0
	for index := 1; index <= sameCount; index++ {
		exponent := float64(index)
		result = result * (math.Pow(float64(a), exponent) + math.Pow(float64(b), exponent) + math.Pow(float64(c), exponent))
	}

	return int(result)
}

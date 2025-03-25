package main

// https://school.programmers.co.kr/learn/courses/30/lessons/181936
func solution(number int, n int, m int) int {
	if number % n == 0 && number % m == 0 {
		return 1
	} else {
		return 0
	}
}
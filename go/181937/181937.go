package main

// https://school.programmers.co.kr/learn/courses/30/lessons/181937?language=go
func solution(num int, n int) int {
	if num % n == 0 {
		return 1
	} else {
		return 0
	}
}
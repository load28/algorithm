package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/181941
func main() {
	data := []string{"a", "b", "c"}
	result := solution(data)

	fmt.Println(result)
}

func solution(arr []string) string {
    var result string

    for _, char := range arr {
    	result += char
    }

    return result
}
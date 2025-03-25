package main

import "fmt"

func main() {
	result := solution("hi", 2)
	fmt.Println(result)
}

func solution(my_string string, k int) string {
	var result string
	for i := 0; i < k; i++ {
		result += my_string
	}
	return result
}

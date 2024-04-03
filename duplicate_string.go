package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "He11oWor1d"
	str2 := "lloWorl"
	index := 2

	result := solution(str1, str2, index)

	fmt.Printf(result)
}


func solution(my_string string, overwrite_string string, s int) string {
	var result string

	str := strings.Split(my_string, "")

	for index, char := range str {
		if index < s {
			result += char
		} else if index < len(overwrite_string) - 1 {
			result += overwrite_string[index]
		} else {
			result += char
		}
	}

	return result
}
package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Program29b8UYP"
	str2 := "merS123"
	index := 7

	result := solution(str1, str2, index)

	fmt.Printf("%s", result)
}

func solution(my_string string, overwrite_string string, s int) string {
	var result string

	str := strings.Split(my_string, "")

	for index, char := range str {
		if index < s || len(overwrite_string)+s <= index {
			result += char
		} else {
			result += string(overwrite_string[index-s])
		}
	}

	return result
}

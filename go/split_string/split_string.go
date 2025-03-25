package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	chars := strings.Split(input, "")
	for _, char := range chars {
		fmt.Println(char)
	}
}
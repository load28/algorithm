package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var result, input string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		break;
	}

	words := strings.Fields(input)
	for _, word := range words {
		result += word
	}


	fmt.Println(result)
}

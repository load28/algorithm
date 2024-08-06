package main

import "fmt"

func main() {
	result := solution("<", "=", 10, 20)
	fmt.Println(result)
}

func solution(ineq string, eq string, n int, m int) int {
	if ineq == "<" && eq == "=" {
		return toInt(n <= m)
	} else if ineq == ">" && eq == "=" {
		return toInt(n >= m)
	} else if ineq == ">" && eq == "!" {
		return toInt(n > m)
	} else {
		return toInt(n < m)
	}
}

func toInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

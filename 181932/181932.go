package main

import "fmt"

func main() {
	result := solution("eidi1fo1e")
	result2 := solution2("eidi1fo1e")
	fmt.Println(result)
	fmt.Println(result2)
}

func solution(code string) string {
	mode := 0
	ret := ""

	for index, char := range code {
		isEven := index%2 == 0
		if string(char) != "1" {
			if mode == 0 {
				if isEven {
					ret += string(char)
				}
			} else {
				if !isEven {
					ret += string(char)
				}
			}
		} else {
			if mode == 0 {
				mode = 1
			} else {
				mode = 0
			}
		}
	}

	if ret == "" {
		return "EMPTY"
	} else {
		return ret
	}
}

// mode 변수를 사용하여 코드를 더 간결하게 작성할 수 있습니다.
func solution2(code string) string {
	mode := 0
	ret := ""

	for index, char := range code {
		isEven := index%2 == 0
		if string(char) != "1" {
			// 뎁스는 최대한 줄이는게 좋다
			if (mode == 0 && isEven) || (mode == 1 && !isEven) {
				ret += string(char)
			}
		} else {
			// 아주 간단한 아이디어지만 좋다
			mode = 1 - mode
		}
	}

	if ret == "" {
		return "EMPTY"
	}
	return ret
}

package main

import "fmt"
import "unicode"
import "strings"

/*
영어 알파벳으로 이루어진 문자열 str이 주어집니다. 
각 알파벳을 대문자는 소문자로 소문자는 대문자로 변환해서 출력하는 코드를 작성해 보세요.
*/
func main() {
	var str string
	var result string

	fmt.Scan(&str)

	for _, runeValue := range str {
		if unicode.IsUpper(runeValue) {
			result += strings.ToLower(string(runeValue))
		} else {
			result += strings.ToUpper(string(runeValue))
		}
	}

	fmt.Println(result)
}
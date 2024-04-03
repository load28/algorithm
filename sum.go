package main

import "fmt"
import "unicode/utf8"

func getStrLen(str string) int {
	return utf8.RuneCountInString(str)
}

func getStrByIndex(str string, index int) string {
	start := 0
	// 문자열은 포인터와 사이즈로 구성되어 있으며 각 사이즈를 가져와서 
	for i := 0; i < index; {
		_, size := utf8.DecodeRuneInString(str[start:])
		start += size
		i++
	}w
	
	r, _ := utf8.DecodeRuneInString(str[start:])
	return runeToString(r)
}

func runeToString(r rune) string {
	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	return string(buf[:n])
}

func main() {
	var s1 string = "안녕하세요 한글"
	r := getStrByIndex(s1, 3)

	fmt.Println(r)
}

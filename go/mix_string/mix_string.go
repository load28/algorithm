package main
 
import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/181942
func main() {
  result := solution("aaa", "bbb")
  fmt.Println(result)
}

// 완성된 문자열의 각 인덱스에 데이터를 넣는 방법이 아닌, 한번에 데이터를 넣는 방법...
func solution(str1 string, str2 string) string {
  var newString string
  for i := 0; i < len(str1); i ++ {
    newString += string(str1[i]) + string(str2[i])
  }

  return newString
}
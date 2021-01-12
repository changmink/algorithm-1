package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*

2. 단어 글자수 세기

제약 조건: 입력 받은 글자는 출력에 포함. 출력은 하나의 문장. 내장 함수로 카운팅할 것

필요한것 : 입출력, 문자열 카운팅 함수

*/
func main() {
	fmt.Println("Input word: ")
	reader := bufio.NewReader(os.Stdin)
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("input Error!")
		return
	}
	word = strings.Replace(word, "\n", "", -1)

	count := len(word)

	fmt.Printf("%s is %d length\n", word, count)
}

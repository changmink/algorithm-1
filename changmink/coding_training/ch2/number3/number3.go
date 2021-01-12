package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	3. 따옴표 출력

	제약 조건: 한개의 출력문으로 결과 출력. String Concatenation 사용(interpolation, subsittution 불가)

	필요한 것: 입출력, 이스케이프 시퀸스
*/

func main() {
	fmt.Print("What is the quote?")
	reader := bufio.NewReader(os.Stdin)
	quote, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	quote = strings.Replace(quote, "\n", "", -1)

	fmt.Print("Who said it?")
	person, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	person = strings.Replace(person, "\n", "", -1)

	var sb strings.Builder
	sb.WriteString(person)
	sb.WriteString(" said. ")
	sb.WriteString("\"")
	sb.WriteString(quote)
	sb.WriteString("\"")

	fmt.Println(sb.String())
}

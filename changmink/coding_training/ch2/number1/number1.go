package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	1. 이름을 입력 받아 인사합니다.

	필요한 것 : stdin, stdout

	제약 조건 : 입력 받는 부분, 스트링 합치는 부분, 출력하는 부분을 따로 작성 할것!
*/

func main() {
	name, err := inputName()
	if err != nil {
		fmt.Println("Stdin error")
		return
	}

	hello := makeHello(name)

	outputHello(hello)
}

func inputName() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name?")
	name, err := reader.ReadString('\n')
	return name, err
}

func makeHello(name string) string {
	hello := fmt.Sprintf("Hello! %s", name)
	return hello
}

func outputHello(hello string) {
	fmt.Println(hello)
}

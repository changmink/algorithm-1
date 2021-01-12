package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	4. Mad Lib

	제약 조건 : 한개의 출력문. interpolation, subsittution 있다면 사용
	필요한 것 : 입출력, 문자열 합성

*/
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a noun: ")
	noun, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	noun = strings.Replace(noun, "\n", "", -1)

	fmt.Print("Enter a verb: ")
	verb, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	verb = strings.Replace(verb, "\n", "", -1)

	fmt.Print("Enter a adjective: ")
	adjective, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	adjective = strings.Replace(adjective, "\n", "", -1)

	fmt.Print("Enter a adverb: ")
	adverb, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	adverb = strings.Replace(adverb, "\n", "", -1)

	fmt.Printf("%s %s %s %s\n", noun, verb, adjective, adverb)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	5. 간단한 수학
	제약 조건 : 문자열로 입력받기, 한개의 출력문만 사용, 입력 값과 출력값은 다른 프로세스에 영향을 받지 않는다.

*/
func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("First number?")
	first, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	first = strings.Replace(first, "\n", "", -1)

	fmt.Print("Second number?")
	second, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	second = strings.Replace(second, "\n", "", -1)

	firstNum, err := strconv.Atoi(first)
	if err != nil {
		fmt.Println(err)
		return
	}

	secondNum, err := strconv.Atoi(second)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s + %s = %d\n%s - %s = %d\n%s * %s = %d\n%s / %s = %d\n",
		first, second, firstNum+secondNum, first, second, firstNum-secondNum,
		first, second, firstNum*secondNum, first, second, firstNum/secondNum)
}

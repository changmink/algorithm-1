package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
	6.퇴직 계산기
	제약 조건 : 입력 값을 숫자로 변환, 올해 년도를 프로그램에 직접 넣지 말 것, 시스템 날짜를 구해서 사용할 것
*/
func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your current age? ")
	currentAge, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("At What age would you like to retire? ")
	retireAge, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
	}

	leftAge := retireAge - currentAge
	fmt.Printf("You have %d years left until you can retire\n", leftAge)

	year := time.Now().Year()
	fmt.Printf("It's %d, so you can retire in %d\n", year, year+leftAge)
}

func getNumber(reader *bufio.Reader) (int, error) {
	numberString, err := reader.ReadString('\n')
	if err != nil {
		return -1, err
	}
	numberString = strings.Replace(numberString, "\n", "", -1)

	number, err := strconv.Atoi(numberString)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return number, nil
}

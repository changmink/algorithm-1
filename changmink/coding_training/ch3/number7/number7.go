package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	7. 직사각형 방의 면적
	제약 조건 : 출력문과 계산 부분을 분리 할 것, 상수를 사용하여 변환 상수를 저장할 것
*/
func main() {
	const constant = 0.09290304
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the length of the room in feet?")
	length, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("What is the width of the room in feet?")
	width, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	squareFeet := width * length
	squareMeter := float64(squareFeet) * constant

	fmt.Printf("The area is %d square feet, %f square methers\n", squareFeet, squareMeter)

}

func getNumber(reader *bufio.Reader) (int, error) {
	numberString, err := reader.ReadString('\n')
	if err != nil {
		return -1, err
	}
	numberString = strings.Replace(numberString, "\n", "", -1)

	number, err := strconv.Atoi(numberString)
	if err != nil {
		return -1, err
	}

	return number, nil
}

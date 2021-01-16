package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
	9. 페인트 계산기
	제약 조건 : 상수를 사용하여 변환 상수를 저장 하기, 올림해서 정수 단위로 구하기
*/
func main() {
	const painted = 9
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the length?")
	length, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("What is the width?")
	width, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	square := length * width
	div := float64(square) / float64(painted)
	paintNeed := math.Ceil(div)
	fmt.Printf("You need %d liters of paint\n", int(paintNeed))
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

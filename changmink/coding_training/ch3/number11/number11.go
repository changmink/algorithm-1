package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	11. 환율 변환
*/

func main() {
	const rateTo = 100.0
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many euros are you exchanging?")
	amountFrom, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("What is exchange rate?")
	rateFrom, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	amountTo := float64(amountFrom*rateFrom) / rateTo
	fmt.Printf("%d Euros at an exchange rate of %d is %0.2f dollars\n", amountFrom, rateFrom, amountTo)
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

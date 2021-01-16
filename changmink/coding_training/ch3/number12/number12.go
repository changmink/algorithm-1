package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	12. 단리 계산
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the principal: ")
	principal, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter the rate of interest: ")
	interest, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter the number of year: ")
	year, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	worth := principal * (1 + interest*year)

	fmt.Printf("The investment will be worth $%d\n", worth)
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

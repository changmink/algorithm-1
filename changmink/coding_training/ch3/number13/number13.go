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
	12. 복리 계산
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

	fmt.Print("Enter the number of times the intereset is compound per year: ")
	n, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	worth := float64(principal) * math.Pow(float64(1+interest/n), float64(n*year))

	fmt.Printf("The investment will be worth $%f\n", worth)
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

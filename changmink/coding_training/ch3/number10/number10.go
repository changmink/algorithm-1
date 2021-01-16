package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const texValue = 0.055
	reader := bufio.NewReader(os.Stdin)

	sum := 0
	for i := 1; i <= 3; i++ {
		fmt.Printf("Price of item %d:", i)
		price, err := getNumber(reader)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Quantity of item %d:", i)
		quantity, err := getNumber(reader)
		if err != nil {
			fmt.Println(err)
			return
		}

		sum += price * quantity
	}

	fmt.Printf("Subtotal: $%d\n", sum)
	fmt.Printf("Tax: $%.2f\n", float64(sum)*texValue)
	fmt.Printf("Total: $%.2f\n", float64(sum)+float64(sum)*texValue)
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

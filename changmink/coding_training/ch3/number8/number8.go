package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	8. 피자 파티
	제약 조건: 조각 개수는 짝수여야 한다.
*/
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many people?")
	people, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("How many pizza?")
	pizza, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("How many pieces are in a pizza?")
	pieceOfPizza, err := getNumber(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	if pieceOfPizza%2 != 0 {
		fmt.Println("Piece should be even")
		return
	}

	allOfPieces := pizza * pieceOfPizza

	eachOfPizza := allOfPieces / people
	leftOfPizza := allOfPieces % people

	fmt.Printf("%d people with %d pizzas\n", people, pizza)
	fmt.Printf("Each person gets %d pieces of pizza\n", eachOfPizza)
	fmt.Printf("There are %d leftover pieces\n", leftOfPizza)
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

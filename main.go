package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(question string, r *bufio.Reader) (string, error) {
	// prompt user for input
	fmt.Print(question)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func checkNumber(num string) bool {
	_, err := strconv.ParseFloat(num, 64)

	if err != nil {
		return false
	} else {
		return true
	}
}

type Method struct {
	plus  string
	minus string
}

func calculate(f float64, s float64, method Method) float64 {
	var result float64

	if method.plus == "plus" {
		result = f + s
	} else if method.minus == "minus" {
		result = f - s
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	allowedMethods := []string{"plus", "minus"}
	firstQuestion, _ := getInput("Enter the first number: ", reader)

	// Check the first question
	if !checkNumber(firstQuestion) {
		fmt.Printf("\"%v\" is not a number\n", firstQuestion)
		os.Exit(1)
	}

	secondQuestion, _ := getInput("Enter the second number: ", reader)

	// Check the other question
	if !checkNumber(secondQuestion) {
		fmt.Printf("\"%v\" is not a number\n", secondQuestion)
		os.Exit(1)
	}

	method, _ := getInput("Enter the method: ", reader)
	methodValid := false

	for _, m := range allowedMethods {
		if method == m {
			methodValid = true
		}
	}

	if !methodValid {
		fmt.Printf("Please use one of \"%v\" methods\n", allowedMethods)
		os.Exit(1)
	}

	f, _ := strconv.ParseFloat(firstQuestion, 64)
	s, _ := strconv.ParseFloat(secondQuestion, 64)

	result := calculate(f, s, Method{plus: method, minus: ""})

	fmt.Println(result)
}

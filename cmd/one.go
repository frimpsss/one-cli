package cmd

import (
	"fmt"
	"strconv"
)

func Add(first string, second string) (result string) {
	n1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("The first number is invalid")
		return
	}

	n2, err := strconv.ParseFloat(second, 64)

	if err != nil {
		fmt.Println("The second number is invalid")
		return
	}

	return fmt.Sprintf("%f", n1+n2)

}

func Subtact(first string, second string) (result string) {
	n1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("The first number is invalid")
		return
	}

	n2, err := strconv.ParseFloat(second, 64)

	if err != nil {
		fmt.Println("The second number is invalid")
		return
	}

	return fmt.Sprintf("%f", n1-n2)

}

func Multiply(first string, second string, shouldRoundUp bool) (result string) {
	n1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("The first number is invalid")
		return
	}

	n2, err := strconv.ParseFloat(second, 64)

	if err != nil {
		fmt.Println("The second number is invalid")
		return
	}
	if shouldRoundUp {
		return fmt.Sprintf("%2.f", n1*n2)
	}

	return fmt.Sprintf("%f", n1*n2)
}

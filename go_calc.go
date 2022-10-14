package main

import (
	"fmt"
	"math"
)

type Input struct {
	a, b float64
	op   string
}

func getInput() (Input, error) {
	var (
		a, b float64
		op   string
	)

	// Getting user input
	fmt.Print("Enter the first number: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		return Input{}, fmt.Errorf("invalid input %w", err)
	}

	fmt.Print("Enter the operator: ")
	_, err = fmt.Scan(&op)
	if err != nil {
		return Input{}, fmt.Errorf("invalid input %w", err)
	}

	if op == "sqrt" || op == "√" {
		return Input{a, 0, op}, nil
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		return Input{}, fmt.Errorf("invalid input %w", err)
	}

	return Input{a, b, op}, nil
}

func calc(a float64, b float64, op string) (float64, error) {
	switch op {
	case "add", "+":
		return a + b, nil
	case "subtract", "-":
		return a - b, nil
	case "divide", "/":
		return a / b, nil
	case "multiply", "*":
		return a * b, nil
	case "pow", "**":
		return math.Pow(a, b), nil
	case "sqrt", "√":
		return math.Sqrt(a), nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", op)
	}
}

func isInt(num float64) bool {
	_, frac := math.Modf(num)
	return frac == 0
}

func main() {
	in, err := getInput()
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}
	a, b, op := in.a, in.b, in.op

	// Calculating result
	res, err := calc(a, b, op)
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	// Checking if the result is integer
	resStr := fmt.Sprintf("%.2f", res)
	if isInt(res) {
		var intRes = int(res)
		resStr = fmt.Sprintf("%d", intRes)
	}

	fmt.Println("Result:", resStr)
}

package main

import (
	"fmt"
	"math"
)

type Input struct {
	a, b float64
	op   string
}

func getInput() Input {
	var (
		a, b float64
		op   string
	)

	// Getting user input
	_, err := fmt.Scan(&a, &op)
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	if op == "sqrt" || op == "√" {
		return Input{a, 0, op}
	}

	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}
	return Input{a, b, op}
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
	in := getInput()
	a, b, op := in.a, in.b, in.op

	// Calculating result
	res, err := calc(a, b, op)
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	// Checking if the result is integer
	resStr := fmt.Sprintf("%.2f", res)
	if isInt(res) {
		var intRes int = int(res)
		resStr = fmt.Sprintf("%d", intRes)
	}

	fmt.Println(resStr)
}

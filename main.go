package main

import (
	"fmt"
	"math"
	"strconv"
)

func executeAndPrint(first float64, second float64, operation string) {

	if operation == ">" || operation == "<" || operation == ">=" || operation == "==" || operation == "<=" || operation == "!=" {
		fmt.Print("RESULT:", compare(first, second, operation))
	} else {
		fmt.Print("RESULT:", calculate(first, second, operation))
	}

}

func compare(first float64, second float64, operation string) (result bool) {

	switch operation {
	case ">":
		result = first > second
	case "<":
		result = first < second
	case "==":
		result = first == second
	case "!=":
		result = first != second

	case ">=":
		result = first >= second
	case "<=":
		result = first <= second
	default:
		fmt.Print("error")
	}

	return result

}

func calculate(first float64, second float64, operation string) (result float64) {

	switch operation {
	case "+":
		result = first + second
	case "-":
		result = first - second
	case "/":
		if second == 0 {
			fmt.Print("error")
		} else {
			result = first / second
		}
	case "%":
		result = math.Remainder(first, second)

	case "*":
		result = first * second
	case "^":
		result = math.Pow(first, second)
	default:
		fmt.Print("error")
	}
	return result

}

func main() {

	var userInput1 string
	var userInput2 string
	var operation string
	var isError bool = false

	fmt.Scanf("%v %v %v", &userInput1, &operation, &userInput2)

	number1, floater1 := strconv.ParseFloat(userInput1, 64)
	number2, floater2 := strconv.ParseFloat(userInput2, 64)

	if operation == "" {
		isError = true
	}

	if int1, er1 := strconv.Atoi(userInput1); er1 == nil {
		fmt.Printf("number1 Type: %T\n", int1)
	} else if floater1 == nil {
		fmt.Printf("number2 Type: %T\n", number1)
	} else {
		fmt.Println("It's not a number (number1)")
		isError = true
	}

	if int2, er2 := strconv.Atoi(userInput2); er2 == nil {
		fmt.Printf("number2 Type: %T\n", int2)
	} else if floater2 == nil {
		fmt.Printf("number2 Type: %T\n", number2)
	} else {
		fmt.Println("It's not a number (number2)")
		isError = true
	}

	if !isError {
		executeAndPrint(number1, number2, operation)
	}

}

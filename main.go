package main

import (
	"fmt"
	"math"
)

func main() {
	var firstNumber float64
	var secondNumber float64
	var operator string

	fmt.Println("Enter the first number: ")
	_, err := fmt.Scanln(&firstNumber)

	if err != nil {
		fmt.Print("The incorrect input of the first number!")
		return 
	}

	fmt.Println("Enter the operation on two numbers: ")
	fmt.Scanln(&operator)

	fmt.Println("Enter the second number: ")
	_, err = fmt.Scanln(&secondNumber)

	if err != nil {
		fmt.Print("The incorrect input of the second number!")
		return 
	}

	switch operator {
	case "+":
		fmt.Println("Answer is", firstNumber + secondNumber)
	case "-":
		fmt.Println("Answer is", firstNumber - secondNumber)
	case "*": 
		fmt.Println("Answer is", firstNumber * secondNumber)
	case "/":
		if (secondNumber == 0) {
			fmt.Println("Division by zero error!")
			return
		}
		fmt.Println("Answer is", firstNumber / secondNumber)
	case "^":
		fmt.Println("Answer is", math.Pow(firstNumber, secondNumber))
	case "%":
		fmt.Println("Answer is", int(firstNumber) % int(secondNumber))
	default: 
		fmt.Println("Cannot compute the result. Not a valid operator - ", operator)
	}
}
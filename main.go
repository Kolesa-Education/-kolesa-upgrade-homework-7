package main

import "fmt"

func main() {
	var (
		operator                          string
		firstNumber, secondNumber, result float64
	)

	fmt.Print("Enter first number: ")
	_, err := fmt.Scanln(&firstNumber)
	if err != nil {
		fmt.Print("Not a number")
		return
	}

	fmt.Print("Enter operator: ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Print("Not valid operator")
		return
	}

	fmt.Print("Enter second number: ")
	_, err = fmt.Scanln(&secondNumber)
	if err != nil {
		fmt.Print("Not a number")
		return
	}

	switch operator {
	case "+":
		result = firstNumber + secondNumber
		break
	case "-":
		result = firstNumber - secondNumber
		break
	case "*":
		result = firstNumber * secondNumber
		break
	case "/":
		if secondNumber == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = firstNumber / secondNumber
		break
	case "%":
		result = firstNumber * (secondNumber / 100)
		break
	default:
		fmt.Println("Invalid operation")
	}
	fmt.Print("Result = ", result)
}

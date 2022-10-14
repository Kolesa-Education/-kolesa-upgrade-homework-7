package main

 import (
 	"fmt"
 	"math"
 )

type Calculator struct{
	firstVariable, secondVariable float64
	operation string
}

func getCalculatorOperatorInput(inputText string) (string, error) {
	var operator string
	fmt.Print(inputText)
	_, err := fmt.Scan(&operator)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	} else {
		return operator, nil
	}
}

func getCalculatorNumericInput(inputText string) (float64, error) {
	var variable float64
	fmt.Print(inputText)
	_, err := fmt.Scan(&variable)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	} else {
		return variable, nil
	}
}

func getCalculatorInputs() (Calculator) {
	firstVariable, err1 := getCalculatorNumericInput("Enter the first variable: ")
	secondVariable, err2 := getCalculatorNumericInput("Enter the second variable: ")
	operator, err3 := getCalculatorOperatorInput("Enter operator: ")
	if err1 == nil && err2 == nil && err3 == nil {
		return Calculator{firstVariable, secondVariable, operator}
	} else {
		fmt.Print("Something went wrong! Please start again!\n")
		return getCalculatorInputs()
	}
}

func calculateAnswer(calculatorVariables Calculator) (float64, error) {
	firstVariable, secondVariable, operation := calculatorVariables.firstVariable, calculatorVariables.secondVariable, calculatorVariables.operation
	switch operation {
	case "+":
		return firstVariable + secondVariable, nil
	case "-":
		return firstVariable - secondVariable, nil
	case "/":
		return firstVariable / secondVariable, nil
	case "*":
		return firstVariable * secondVariable, nil
	case "**":
		return math.Pow(firstVariable, secondVariable), nil
	default:
		return 0, fmt.Errorf(operation)
	}
}

 func main() {
	calculatorVariables := getCalculatorInputs()
	calculatedAnswer, err := calculateAnswer(calculatorVariables)
	if err != nil {
		fmt.Print("Something went wrong! Please start again!\n", err)
	} else {
		fmt.Print(calculatedAnswer)
	}
 }
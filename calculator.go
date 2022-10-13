package main

import (
	"fmt"
	"math"
)

func main() {
	var number1, number2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&number1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&number2)

	fmt.Print("Enter the operator +, -, *, /, **: ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%.1f %s %.1f = %.1f", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%.1f %s %.1f = %.1f", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%.1f %s %.1f = %.1f", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Println("Can't divide by zero")
		} else {
			fmt.Printf("%.1f %s %.1f = %.1f", number1, operator, number2, number1/number2)
		}

	case "**":
		result := math.Pow(float64(number1), float64(number2))
		fmt.Printf("%.1f %s %.1f = %.1f", number1, operator, number2, result)

	default:
		fmt.Println("Invalid operator!")
	}

}

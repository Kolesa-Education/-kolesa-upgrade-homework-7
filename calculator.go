package main

import (
	"fmt"
)

func main() {
	var number1, number2 float64
	var operator string

	fmt.Print("Enter the first number : ")
	fmt.Scanln(&number1)

	fmt.Print("Enter the second number : ")
	fmt.Scanln(&number2)

	fmt.Print("Enter the Operator (+ - * /) : ")
	fmt.Scanln(&operator)

	switch operator {

	case "+":
		fmt.Println(number1, operator, number2, "=", number1+number2)

	case "-":
		fmt.Println(number1, operator, number2, "=", number1-number2)

	case "*":
		fmt.Println(number1, operator, number2, "=", number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Println("Can not divide by 0")
		} else {
			fmt.Println(number1, operator, number2, "=", number1/number2)
		}

	default:
		fmt.Println("Wrong operator, put correct values everywhere")
	}

}

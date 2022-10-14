package main

import (
	"fmt"
	"math"
)

func main() {
	var number1, number2 float64
	var operator string

	fmt.Print("Enter the first number : ")
	_, err := fmt.Scanln(&number1)
	if err != nil {
		fmt.Println("Not a number, please enter number for the first value")
		return
	}

	fmt.Print("Enter the second number : ")
	_, err = fmt.Scanln(&number2)
	if err != nil {
		fmt.Print("Not a number, please enter number for the second value")
		return
	}

	fmt.Print("Enter the Operator (+ - * / > < = ^ %) : ")
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

	case ">":
		if number1 < number2 {
			fmt.Println("False, ", number1, "<", number2)
		} else {
			fmt.Println("True, ", number1, ">", number2)
		}
	case "<":
		if number1 > number2 {
			fmt.Println("False, ", number1, ">", number2)
		} else {
			fmt.Println("True, ", number1, "<", number2)
		}

	case "=":
		if number1 == number2 {
			fmt.Println("True, ", number1, "=", number2)
		} else {
			fmt.Println("False, ", number1, "≠", number2)

		}

	case "^":
		fmt.Printf("%.2f", math.Pow(number1, number2))
		break

	case "%":
		fmt.Printf("%.2f", float64(int(number1)%int(number2)))
		break

	default:
		fmt.Println("Wrong operator, put correct values everywhere")
	}

}

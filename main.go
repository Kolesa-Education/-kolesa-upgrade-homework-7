package main

import (
	"fmt"
	"math"
)

func main() {
	var number1 float64
	var number2 float64
	var operand string

	fmt.Println("Введите первое число")
	_, err := fmt.Scanln(&number1)

	if err != nil {
		fmt.Print("Введите правильно первое число")
		return
	}

	fmt.Println("Введите операнд")
	fmt.Scanln(&operand)

	fmt.Println("Введите второе число")
	_, err = fmt.Scanln(&number2)

	if err != nil {
		fmt.Print("Введите правильно второе число")
		return
	}

	switch operand {
	case "+":
		fmt.Printf("%v %s %v = %v", number1, operand, number2, addition(number1, number2))
	case "/":
		if number2 == 0 {
			fmt.Printf("На ноль делить нельзя")
			return
		}
		fmt.Printf("%v %s %v = %v", number1, operand, number2, divide(number1, number2))
	case "*":
		fmt.Printf("%v %s %v = %v", number1, operand, number2, multiplication(number1, number2))
	case "%":
		fmt.Printf("%v %s %v = %v", number1, operand, number2, percent(number1, number2))
	case "mod":
		fmt.Printf("%v %s %v = %v", number1, operand, number2, mod(int(number1), int(number2)))
	case "-":
		fmt.Printf("%v %s %v = %v", number1, operand, number2, subtraction(number1, number2))
	case "^":
		fmt.Printf("%v %s %v = %v", number1, operand, number2, pow(number1, number2))
	default:
		fmt.Println("Такого оператора не существует")
	}

}

func addition(number1, number2 float64) int {
	return int(number1 + number2)
}

func divide(number1, number2 float64) float64 {
	return number1 / number2
}

func multiplication(number1, number2 float64) int {
	return int(number1 * number2)
}

func subtraction(number1, number2 float64) int {
	return int(number1 - number2)
}

func mod(number1, number2 int) int {

	return number1 % number2
}

func percent(number1, number2 float64) int {
	return int((number1 * number2) / 100)
}

func pow(number1, number2 float64) float64 {
	return math.Pow(number1, number2)
}

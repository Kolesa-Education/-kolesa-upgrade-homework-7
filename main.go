package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var firstOperand, operator, secondOperand string
	fmt.Println("Введите выражение (по примеру: 2 + 3):")
	fmt.Println("Допустимые операторы: +, -, *, /, %, ^, p (процент)")
	fmt.Scanln(&firstOperand, &operator, &secondOperand)
	Calc(firstOperand, operator, secondOperand)
}

func Calc(firstOperand, operator, secondOperand string) {
	firstNumber, err := convertToNumber(firstOperand)
	if err != nil {
		fmt.Println("Ошибка: первый операнд не является числом или не назначен")
		return
	}
	secondNumber, err := convertToNumber(secondOperand)
	if err != nil {
		fmt.Println("Ошибка: второй операнд не является числом или не назначен")
		return
	}
	switch operator {
	case "+":
		fmt.Print(firstNumber + secondNumber)
	case "-":
		fmt.Print(firstNumber - secondNumber)
	case "*":
		fmt.Print(firstNumber * secondNumber)
	case "/":
		if secondNumber == 0 {
			fmt.Print("Нельзя делить на 0 (ноль)")
			return
		}
		fmt.Print(firstNumber / secondNumber)
	case "%":
		if secondNumber == 0 {
			fmt.Print("Нельзя делить на 0 (ноль) по модулю")
			return
		}
		fmt.Print(math.Mod(firstNumber, secondNumber))
	case "^":
		fmt.Print(math.Pow(firstNumber, secondNumber))
	case "p":
		fmt.Print(firstNumber / 100 * secondNumber)
	default:
		fmt.Printf("Неверный оператор: %s", operator)
		return
	}
}

func convertToNumber(num string) (float64, error) {
	resFloat, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, err
	}
	return resFloat, nil
}

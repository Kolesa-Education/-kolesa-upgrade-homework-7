package main

import (
	"fmt"
	"math"
)

func main() {
	var number1, number2 int
	var operator string

	fmt.Print("Введите первое число : ")
	fmt.Scanln(&number1)

	fmt.Print("Введите оператор  | + | - | * | / | ^ | √ | <=> | : ")
	fmt.Scanln(&operator)

	fmt.Print("Введите второе число : ")
	fmt.Scanln(&number2)

	switch operator {
	case "+":
		fmt.Printf("%d %s %d = %d", number1, operator, number2, number1+number2)
	case "-":
		fmt.Printf("%d %s %d = %d", number1, operator, number2, number1-number2)
	case "*":
		fmt.Printf("%d %s %d = %d", number1, operator, number2, number1*number2)
	case "^":
		fmt.Printf("%d %s %d = %.2f", number1, operator, number2, math.Pow(float64(number1), float64(number2)))
	case "/":
		if number2 == 0 {
			fmt.Println("Проверьте правильность введенных данных")
			fmt.Println("Нельзя делить на 0")
		} else {
			fmt.Printf("%d %s %d = %d", number1, operator, number2, number1/number2)
		}
	case "√":
		if number2 == 0 {
			fmt.Println("Проверьте правильность введенных данных")
			fmt.Println("Невозможно извлечь корень 0-ой степени")
		} else {
			fmt.Printf("%d %s %d = %.2f", number1, operator, number2, math.Pow(float64(number1), 1/float64(number2)))
		}

	case "<=>":
		if number1 == number2 {
			fmt.Printf("%d = %d", number1, number2)
		} else if number1 > number2 {
			fmt.Printf("%d > %d", number1, number2)
		} else {
			fmt.Printf("%d < %d", number1, number2)
		}

	default:
		fmt.Println("Проверьте правильность введенных данных")
		fmt.Println("Неправильно указан оператор или операнд")
	}

}

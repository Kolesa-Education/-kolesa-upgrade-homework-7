package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64
	var num2 float64
	var operator string

	fmt.Println("Введите операнд 1")
	_, err := fmt.Scanln(&num1)

	if err != nil {
		fmt.Print("Операнды должны быть численными")
		return
	}

	fmt.Println("Введите операнд 2")
	_, err = fmt.Scanln(&num2)

	if err != nil {
		fmt.Print("Операнды должны быть численными")
		return
	}

	fmt.Println("Введите операцию")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("Результат: %v", num1+num2)
	case "-":
		fmt.Printf("Результат: %v", num1-num2)
	case "/":
		if num2 == 0 {
			fmt.Printf("Нельзя делить на ноль")
			return
		}
		fmt.Printf("Результат: %v", num1/num2)
	case "*":
		fmt.Printf("Результат: %v", num1*num2)
	case "^":
		fmt.Printf("Результат: %v", math.Pow(num1, num2))
	case "%":
		fmt.Printf("Результат: %v", num1/100*num2)
	case "mod":
		fmt.Printf("Результат: %v", int(num1)%int(num2))
	default:
		fmt.Println("значение c=%s' не является оператором", operator)
	}

}

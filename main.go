package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	var op string

	fmt.Println("Введите первое число: ")
	_, err := fmt.Scanln(&num1)

	if err != nil {
		fmt.Print("Please, enter a number")
		return
	}

	fmt.Println("Введите второе число: ")
	_, err = fmt.Scanln(&num2)

	if err != nil {
		fmt.Print("Please, enter a number")
		return
	}

	fmt.Println("Введите операцию: ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		fmt.Printf("%f %s %f = %f", num1, op, num2, num1+num2)
	case "-":
		fmt.Printf("%f %s %f = %f", num1, op, num2, num1-num2)
	case "*":
		fmt.Printf("%f %s %f = %f", num1, op, num2, num1*num2)
	case "%":
		fmt.Printf("%f %s %f = %f", num1, op, num2, (num1/num2)*100)
	case "/":
		if num2 == 0 {
			fmt.Printf("Нельзя делить на ноль")
			return
		}
		fmt.Printf("%f %s %f = %f", num1, op, num2, num1/num2)
	case "^":
		fmt.Printf("%f %s %f = %f", num1, op, num2, math.Pow(num1, num2))
	case "<":
		fmt.Printf("%f %s %f is %v", num1, "<", num2, num1 < num2)
	case ">":
		fmt.Printf("%f %s %f is %v", num1, ">", num2, num1 > num2)
	case "=":
		fmt.Printf("%f %s %f is %v", num1, "=", num2, num1 == num2)
	default:
		fmt.Println("Этот оператор не поддерживается данным калькулятором")
	}

}

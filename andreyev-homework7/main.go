package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Введите первое число :")
	fmt.Scanln(&num1)

	fmt.Print("Введите второе число :")
	fmt.Scanln(&num2)

	fmt.Print("Введите операцию (+ -  * /  ^  <  > ) ")
	fmt.Scanln(&operator)

	switch operator {

	case "+":
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, num1+num2)

	case "-":
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, num1-num2)

	case "*":
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, num1*num2)

	case "/":
		if num2 == 0 {
			fmt.Println("Нельзя делить на 0")
		} else {
			fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, num1/num2)
		}

	case "^":
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, math.Pow(num1, num2))

	case "<":
		if num1 < num2 {
			fmt.Println(num1, "<", num2)
		} else if num1 > num2 {
			fmt.Println(num1, ">", num2)
		} else {
			fmt.Println(num1, "=", num2)
		}

	case ">":
		if num1 > num2 {
			fmt.Println(num1, ">", num2)
		} else if num1 < num2 {
			fmt.Println(num1, "<", num2)
		} else {
			fmt.Println(num1, "=", num2)
		}

	default:
		fmt.Println("Некорректный оператор")
	}

}

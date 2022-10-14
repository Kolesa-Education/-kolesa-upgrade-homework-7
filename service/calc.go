package service

import (
	"fmt"
	"math"
)

func Calc(tempNum1, oper, tempNum2 string) {
	num1, err := convertFloat(tempNum1)
	if err != nil {
		fmt.Println("Одно из значении не является числом")
		return
	}
	num2, err := convertFloat(tempNum2)
	if err != nil {
		fmt.Println("Одно из значении не является числом")
		return
	}
	switch oper {
	case "+":
		fmt.Println("Ответ: ", num1+num2)
	case "-":
		fmt.Println("Ответ: ", num1-num2)
	case "*":
		fmt.Println("Ответ: ", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("No divide by 0")
			return
		}
		fmt.Println("Ответ: ", num1/num2)
	case "%":
		if num2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		fmt.Println("Ответ: ", math.Mod(num1, num2))
	case "^":
		fmt.Println("Ответ: ", math.Pow(num1, num2))
	default:
		fmt.Printf("%s: Invalid operation", oper)
		return
	}
}

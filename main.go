package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var tempNum1 string
	var oper string
	var tempNum2 string

	fmt.Println("Введите математическое выражение")
	fmt.Scan(&tempNum1, &oper, &tempNum2)

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

	calcManipulation(num1, num2, oper)

}

func calcManipulation(num1, num2 float64, oper string) {
	switch oper {
	case "+":
		fmt.Println("Ответ:", num1+num2)
	case "-":
		fmt.Println("Ответ:", num1-num2)
	case "*":
		fmt.Println("Ответ:", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("No division by 0")
			return
		}
		fmt.Println("Ответ:", num1/num2)
	case "%":
		if num2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		fmt.Println("Ответ:", math.Mod(num1, num2))
	case "^":
		fmt.Println("Ответ:", math.Pow(num1, num2))
	default:
		fmt.Println(oper, ":Invalid operation")
	}
}

func convertFloat(num string) (float64, error) {
	resFloat, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, err
	}
	return resFloat, nil
}

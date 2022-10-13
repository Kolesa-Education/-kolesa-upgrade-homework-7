package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, res float64
	var op string
	fmt.Println("введи два числа и операцию над ними (+, -, *, /, ^), например," +
		" для возведения в степень введи 2^5")
	_, err := fmt.Scan(&a, &op, &b)
	if err != nil {
		fmt.Println("Нужно ввести числа!")
		return
	}
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Делить на ноль нельзя!")
			return
		} else {
			res = a / b
		}
	case "^":
		res = math.Pow(a, b)
	default:
		fmt.Println("Неизвестная операция")
	}
	fmt.Println("Ответ: ", res)
}

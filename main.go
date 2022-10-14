package main

import (
	"fmt"
	"math"
)

func main() {
  var num1 float64
	var num2 float64
	var operator string
  fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter an operator (+, -, *, /, %, ?, ^): ")
	fmt.Scanln(&operator)
  fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)
	calculate(num1, num2, operator)
}

func calculate(num1 float64, num2 float64, operator string) {
	switch operator {
		case "+":
			fmt.Println(num1 + num2)
		case "-":
			fmt.Println(num1 - num2)
		case "*":
			fmt.Println(num1 * num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Division by zero is not allowed!")
			} else {
				fmt.Println(num1 / num2)
			}
		case "%":
			if num2 == 0 {
				fmt.Println("Division by zero is not allowed!")
			} else {
				fmt.Println(math.Mod(num1, num2))
			}
		case "^":
			fmt.Println(math.Pow(num1, num2))
		case "?":
			if num1 > num2 {
				fmt.Println(num1, "is bigger than", num2)
			} else if num1 < num2 {
				fmt.Println(num1, "is less than", num2)
			} else {
				fmt.Println(num1, "is equal to", num2)
			}
		default:
			fmt.Println("Invalid operator! You can choose only one operator from suggested list.")
	}
}

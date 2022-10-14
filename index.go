package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		num1, num2 float64
		oper       string
	)

	fmt.Println("The first number: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Invalid input!!!")
		return
	}

	fmt.Println("The second number: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("Invalid input!!!")
		return
	}

	fmt.Println("Select the operator(+ - / * % ^ < > ==): ")
	fmt.Scanln(&oper)

	switch oper {
	case "+":
		fmt.Print(num1, " + ", num2, " = ")
		fmt.Println(num1 + num2)
	case "-":
		fmt.Print(num1, " - ", num2, " = ")
		fmt.Println(num1 - num2)
	case "*":
		fmt.Print(num1, " * ", num2, " = ")
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("can not divide to 0")
			break
		}
		fmt.Print(num1, " / ", num2, " = ")
		fmt.Println(num1 / num2)
	case "%":
		fmt.Print(num1, " % ", num2, " = ")
		fmt.Println(int(num1) % int(num2))
	case "<":
		fmt.Print(num1, " < ", num2, ": ")
		fmt.Println(num1 < num2)
	case ">":
		fmt.Print(num1, " > ", num2, ": ")
		fmt.Println(num1 > num2)
	case "==":
		fmt.Print(num1, " == ", num2, ": ")
		fmt.Println(num1 == num2)
	case "^":
		fmt.Print(num1, " ^ ", num2, " = ")
		fmt.Println(math.Pow(num1, num2))
	default:
		fmt.Println("Wrong operation!!!")
	}
}

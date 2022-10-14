package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("неправильный ввод\nпример:\n2 + 5")
		return
	}
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("неправильный ввод\nпример:\n2 + 5")
		return
	}

	o := os.Args[2]
	y, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("неправильный ввод\nпример:\n2 + 5")
		return
	}
	fmt.Printf("%d %s %d = ", x, o, y)
	switch o {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		if y == 0 {
			fmt.Println("\nделить на ноль нельзя")
			return
		}
		fmt.Println(x / y)
	case "%":
		if y == 0 {
			fmt.Println("\nделить по модулю на ноль нельзя")
			return
		}
		fmt.Println(x % y)
	default:
		fmt.Println("\nнеправильный ввод:\nвторым пишите оператор")
	}
}

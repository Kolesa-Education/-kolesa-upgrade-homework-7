package main

import (
	"fmt"
)

func main() {
	var c string
	var a, b int
	var output int

	fmt.Print("Введите первое значение: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе значение: ")
	fmt.Scanln(&b)
	fmt.Print("Введите операцию (+,-,/,*,%):")
	fmt.Scanln(&c)

	output = 0

	switch c {
	case "+":
		output = a + b
		break
	case "-":
		output = a - b
	case "*":
		output = a * b
	case "/":
		output = a / b
	case "%":
		output = a % b
		// output = math.Mod(a, b)
	default:
		fmt.Println("Не верная операция")
	}
	fmt.Printf("%d %s %d = %d", a, c, b, output)
}

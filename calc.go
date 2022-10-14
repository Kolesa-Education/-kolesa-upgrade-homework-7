package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b float64
	var operator string

	fmt.Println("Калькулятор")
	fmt.Println("Программа выполняет простые арифметические действия в виде A+B.")
	fmt.Println("Доступные арифметические операции:")
	fmt.Println("+ - сложение")
	fmt.Println("- - вычитание")
	fmt.Println("* - умножение")
	fmt.Println("/ - деление")
	fmt.Println("^ - возведение числа А в ступень B")
	fmt.Println("K - корень из числа А в степени B")
	fmt.Println("% - вычисление процента B от числа A")
	fmt.Println("S - сравнение чисел")

	fmt.Print("Введите число A: ")
	fmt.Scanln(&a)

	fmt.Print("Введите число B: ")
	fmt.Scanln(&b)

	fmt.Print("Выберете оператор: ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f %s %f = %f", a, operator, b, a+b)
	case "-":
		fmt.Printf("%f %s %f = %f", a, operator, b, a-b)
	case "*":
		fmt.Printf("%f %s %f = %f", a, operator, b, a*b)
	case "/":
		if b == 0.0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Printf("%f %s %f = %f", a, operator, b, a/b)
		}
	case "^":
		fmt.Printf("%f %s %f = %f", a, operator, b, math.Pow(a, b))
	case "K":
		if b == 0 {
			fmt.Println("Корень в 0-ой степени извлечь нельзя!")
		} else {
			fmt.Printf("%f %s %f = %f", a, operator, b, math.Pow(a, 1/b))
		}
	case "%":
		if b == 0 {
			fmt.Println("Корень в 0-ой степени извлечь нельзя!")
		} else {
			fmt.Printf("%f %s %f = %f", a, operator, b, a*b/100)
		}
	case "S":
		if a == b {
			fmt.Printf("%f = %f", a, b)
		} else if a > b {
			fmt.Printf("%f > %f", a, b)
		} else {
			fmt.Printf("%f < %f", a, b)
		}
	default:
		fmt.Println("Недопустимый оператор!")
	}
}

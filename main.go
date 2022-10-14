package main

import (
	"fmt"
	"math"
)

func main() {

	var a float64
	var b float64
	var c string

	fmt.Print("A = ")
	fmt.Scanln(&a)
	fmt.Print("B = ")
	fmt.Scanln(&b)
	fmt.Println("Please choose operator from list: + - * / pow < > = %")
	fmt.Scanln(&c)

	switch c {
	case "+":
		fmt.Println(a + b)
		break
	case "-":
		fmt.Println(a - b)
		break
	case "*":
		fmt.Println(a * b)
		break
	case "/":
		if b == 0 {
			fmt.Println("0 divide")
		} else {
			fmt.Println(a / b)
		}
		break
	case ">":
		if a > b {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
		break
	case "<":
		if a < b {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	case "=":
		if a == b {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
		break
	case "%":
		fmt.Println(a * (b / 100))
		break
	case "pow":
		fmt.Println(math.Pow(float64(a), float64(b)))
		break
	default:
		fmt.Println("Invalid operator")
	}
}

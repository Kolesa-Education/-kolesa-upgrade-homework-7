package main

import (
	    "fmt"
	    "math"
       )

func main() {

	var a, b float64
	var opr string


	fmt.Print("Введите первый знаменатель: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второй знаменатель: ")
	fmt.Scanln(&b)
	fmt.Print("Выберите оператор +, -, *, /, **, <, > : ")
	fmt.Scanln(&opr)

	switch opr {
	case "+":
		fmt.Printf("%.1f %s %.1f = %.1f", a, opr, b, a+b)
	case "-":
		fmt.Printf("%.1f %s %.1f = %.1f", a, opr, b, a-b)

	case "*":
		fmt.Printf("%.1f %s %.1f = %.1f", a, opr, b, a*b)

	case "/":
		if b == 0.0 {
			fmt.Println("Нельзя делить на ноль!")
		} else {
			fmt.Printf("%.1f %s %.1f = %.1f", a, opr, b, a/b)
		}

	case "**":
		res := math.Pow(float64(a), float64(b))
		fmt.Printf("%.1f %s %.1f = %.1f", a, opr, b, res)

	case "<":
		if a < b {
			fmt.Println(a, "<", b)
		} else if a > b {
			fmt.Println(a, ">", b)
		} else {
			fmt.Println(a, "=", b)
		}

	case ">":
		if a > b {
			fmt.Println(a, ">", b)
		} else if a < b {
			fmt.Println(a, "<", b)
		} else {
			fmt.Println(a, "=", b)
		}


	default:
		fmt.Println("Неизвестный оператор!")
	}

}
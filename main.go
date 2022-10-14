package main

import "ftm"

func main() {
    var num1, num2 float64
    var operator string

    fmt.Print("Введите первое число: ")
    fmt.Scanln(&num1)

    fmt.Print("Введите второе число: ")
    fmt.Scanln(&num2)

    fmt.Print("Введите оператор (+, -, *, /): ")
    fmt.Scanln(&operator)

    switch operator {
    case "+":
    	fmt.Printf("%.1f %s %.1f = %.1f", num1, operator, num2, num1+num2)

    case "-":
    	fmt.Printf("%.1f %s %.1f = %.1f", num1, operator, num2, num1-num2)

    case "*":
    	fmt.Printf("%.1f %s %.1f = %.1f", num1, operator, num2, num1*num2)

    case "/":
    	if num2 == 0.0 {
    		fmt.Println("Нельзя делить на ноль")
    	} else {
    		fmt.Printf("%.1f %s %.1f = %.1f", num1, operator, num2, num1/num2)
    	}

    default:
    	fmt.Println("Неверный оператор!")
    }
}
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var operator string
	var strnum1, strnum2 string
	var num1, num2 float64
	var floatOutput float64
	fmt.Print("Введите первое число ")
	fmt.Scanf("%s\n", &strnum1)
	fmt.Print("Введите второе число ")
	fmt.Scanf("%s\n", &strnum2)
	fmt.Print("Введите оператор (+,-,*,/,%,pow,>,<,=): ")
	fmt.Scanf("%s\n", &operator)

	//Блин как же не привычно с динамического пыха на статический го переходить, но привыкнуть думаю можно))
	num1, err := strconv.ParseFloat(strnum1, 64)
	if err != nil {
		fmt.Printf("'%s' - это не число", strnum1)
		return
	}

	num2, err = strconv.ParseFloat(strnum2, 64)

	if err != nil {
		fmt.Printf("'%s' - это не число", strnum2)
		return
	}
	switch operator {
	case "+":
		floatOutput = num1 + num2
		fmt.Printf("%.2f", floatOutput)
		break
	case "-":
		floatOutput = num1 - num2
		fmt.Printf("%.2f", floatOutput)
		break
	case "*":
		floatOutput = num1 * num2
		fmt.Printf("%.2f", floatOutput)
		break
	case "/":
		if num2 == 0 {
			fmt.Println("На 0 делить нельзя!")
		} else {
			floatOutput = num1 / num2
		}
		fmt.Printf("%.2f", floatOutput)
		break
	case "%":
		floatOutput = float64(int(num1) % int(num2))
		fmt.Printf("%.2f", floatOutput)
		break
	case "pow":
		floatOutput = math.Pow(num1, num2)
		fmt.Printf("%.2f", floatOutput)
		break
	case ">":
		if num1 > num2 {
			fmt.Printf("Число %.2f больше числа %.2f", num1, num2)
		} else if num1 < num2 {
			fmt.Printf("Число %.2f больше числа %.2f", num2, num1)
		} else if num1 == num2 {
			fmt.Printf("Числа %.2f и %.2f равны", num1, num2)
		}
		break
	case "<":
		if num1 > num2 {
			fmt.Printf("Число %.2f больше числа %.2f", num1, num2)
		} else if num1 < num2 {
			fmt.Printf("Число %.2f больше числа %.2f", num2, num1)
		} else if num1 == num2 {
			fmt.Printf("Числа %.2f и %.2f равны", num1, num2)
		}
		break
	case "=":
		if num1 == num2 {
			fmt.Printf("Числа %.2f и %.2f равны", num1, num2)
		} else {
			fmt.Printf("Числа %.2f и %.2f НЕ равны", num1, num2)
		}
		break

	default:
		fmt.Println("Неверная операция!")
	}
}

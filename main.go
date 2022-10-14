package main

import (
	"fmt"
	"math"
	"strconv"
)

func calculatePrint(first float64, second float64, operation string) {

	fmt.Print("RESULT: ")
	switch operation {
	case "+":
		fmt.Print(first + second)
	case "-":
		fmt.Print(first - second)
	case "/":
		fmt.Print(first / second)
	case "%":
		fmt.Print(math.Remainder(first, second))

	case "*":
		fmt.Print(first * second)
	case "^":
		fmt.Print(math.Pow(first, second))
	case ">":
		fmt.Print(first > second)
	case "<":
		fmt.Print(first > second)
	case "==":
		fmt.Print(first == second)
	case ">=":
		fmt.Print(first >= second)
	case "<=":
		fmt.Print(first >= second)
	default:
		fmt.Print("error")
	}

}

func main() {

	var userInput1 string
	var userInput2 string
	var operation string

	fmt.Scan(&userInput1)
	fmt.Scan(&operation)
	fmt.Scan(&userInput2)

	fmt.Scanf("%v %v %v", &userInput1, &operation, &userInput2)

	if num1, er1 := strconv.Atoi(userInput1); er1 == nil {
		fmt.Printf("number1 = %T\n", num1)

		if num2, er2 := strconv.Atoi(userInput2); er2 == nil {
			fmt.Printf("number2 = %T\n", num2)

			calculatePrint(float64(num1), float64(num2), operation)
		} else {

			if floatnum2, er3 := strconv.ParseFloat(userInput1, 64); er3 == nil {
				fmt.Printf("number2 = %T\n", floatnum2)
				calculatePrint(float64(num1), floatnum2, operation)
			}
		}

	} else {
		if floatnum1, fler3 := strconv.ParseFloat(userInput1, 64); fler3 == nil {
			fmt.Printf("number1 = %T\n", floatnum1)
			if num2, er2 := strconv.Atoi(userInput2); er2 == nil {
				fmt.Printf("number2 = %T\n", num2)
				calculatePrint(floatnum1, float64(num2), operation)
			} else {
				if floatnum2, er3 := strconv.ParseFloat(userInput1, 64); er3 == nil {
					fmt.Printf("number2 = %T\n", floatnum2)

					fmt.Println(floatnum1, floatnum2)
					calculatePrint(floatnum1, floatnum2, operation)
				}
			}
		}

	}

}

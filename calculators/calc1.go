package calculators

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var ErrDivideByZero = errors.New("divide by zero")

func Divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, ErrDivideByZero
	}
	return num1 / num2, nil
}

func Calc1(input string) {
	var num1, num2 float64
	var operator string
	var strnum1, strnum2 string

	parts := strings.Split(input, " ")

	if len(parts) != 3 {
		println("Please enter expression as in example")
		return
	}

	strnum1 = parts[0]
	operator = parts[1]
	strnum2 = parts[2]

	num1, err := strconv.ParseFloat(strnum1, 64)

	if err != nil {
		fmt.Printf("'%s' is not a number", strnum1)
		return
	}

	num2, err = strconv.ParseFloat(strnum2, 64)

	if err != nil {
		fmt.Printf("'%s' is not a number", strnum2)
		return
	}

	/*fmt.Print("Enter the first number: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("\nPlease enter a valid number: ")
		return
	}

	fmt.Print("Enter the operator ( + - * / % ** ): ")
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("\nPlease enter a valid operator: ")
		return
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scan(&input)
	if err != nil {
		fmt.Println("\nPlease enter a valid number: ")
		return
	}*/

	result := 0.0
	switch operator {
	case "+":
		result = num1 + num2
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, result)
	case "-":
		result = num1 - num2
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, result)
	case "*":
		result = num1 * num2
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, result)
	case "/":
		result = num1 / num2
		result, err := Divide(num1, num2)
		if err != nil {
			switch {
			case errors.Is(err, ErrDivideByZero):
				fmt.Println("divide by zero error")
			default:
				fmt.Printf("unexpected division error: %s\n", err)
			}
			return
		}
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, result)

	case "%":
		result = math.Mod(num1, num2)
		result, err := Divide(num1, num2)
		if err != nil {
			switch {
			case errors.Is(err, ErrDivideByZero):
				fmt.Println("mod by zero error")
			default:
				fmt.Printf("unexpected division error: %s\n", err)
			}
			return
		}
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, result)
	case "**":
		result = math.Pow(num1, num2)
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, result)
	default:
		fmt.Println("Invalid operator")
	}
}

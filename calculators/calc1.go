package calculators

import (
	"errors"
	"fmt"
	"math"
)

var ErrDivideByZero = errors.New("divide by zero")

func Divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, ErrDivideByZero
	}
	return num1 / num2, nil
}

func Calc1(text string) {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the operator ( + - * / % **): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	result := 0.0
	switch operator {
	case "+":
		result = num1 + num2
		fmt.Printf("%g %s %g = %g", num1, operator, num2, result)
	case "-":
		result = num1 - num2
		fmt.Printf("%g %s %g = %g", num1, operator, num2, result)
	case "*":
		result = num1 * num2
		fmt.Printf("%g %s %g = %g", num1, operator, num2, result)
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
		fmt.Printf("%g %s %g = %g", num1, operator, num2, result)
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
		fmt.Printf("%g %s %g = %g", num1, operator, num2, result)
	case "**":
		result = math.Pow(num1, num2)
		fmt.Printf("%g %s %g = %g", num1, operator, num2, result)
	default:
		fmt.Println("Invalid operator")
	}
}

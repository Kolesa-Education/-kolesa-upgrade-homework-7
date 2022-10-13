package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var a float64
	var operation string
	var b float64
	fmt.Scan(&a, &operation, &b)
	calc(a, operation, b)
}
func calc(a float64, operation string, b float64) {
	switch operation {
	case string('+'):
		fmt.Println(a + b)
	case string('-'):
		fmt.Println(a - b)
	case string('*'):
		fmt.Println(a * b)
	case string('/'):
		if b == 0 {
			var ErrZeroDivisionAttempt = errors.New("divide by zero is not allowed")
			fmt.Println(ErrZeroDivisionAttempt)
		} else {
			fmt.Println(a / b)
		}
	case string('^'):
		fmt.Println(math.Pow(a, b))
	default:
		fmt.Println("unknown operator: " + operation)
	}
}

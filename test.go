package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

func main() {
	var a int
	var operation string
	var b int
	fmt.Scan(&a, &operation, &b)
	calc(a, operation, b)

}
func calc(a int, operation string, b int) {
	if reflect.TypeOf(a).Kind() != reflect.Int || reflect.TypeOf(operation).Kind() != reflect.String || reflect.TypeOf(b).Kind() != reflect.Int {
		fmt.Println("Err input")
	} else {
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
			fmt.Println(math.Pow(float64(a), float64(b)))
		default:
			fmt.Println("unknown operator: " + operation)
		}
	}
}

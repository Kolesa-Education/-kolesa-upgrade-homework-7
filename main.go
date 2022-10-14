package main

import (
	"fmt"
	"math"
)

func opCalc(a float64, b float64, oper string) float64 {
	switch oper {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		if b == 0 {
			fmt.Println(" division by 0")
			return 0.0
		}
		return a / b
	case "*":
		return a * b
	case "**":
		return math.Pow(a, b)
	case "√":
		return math.Sqrt(a)
	default:
		fmt.Errorf("invalid operator: %s", oper)
		return 0.0
	}
}
func main() {
	var name string
	var a float64
	var b float64

	fmt.Printf("Write first value: ")
	fmt.Scanf("%f", &a)
	fmt.Printf("Write operator : ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Write second value : ")
	fmt.Scanf("%f", &b)
	fmt.Printf("answer  : ")
	fmt.Println(opCalc(a, b, name))
}

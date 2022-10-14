package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Привет, это простой калькулятор")
	var a, b float64
	var operator string

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&operator)
	res, err := Calc(a, b, operator)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch res.(type) {
	case float64:
		fmt.Println(res)
	case bool:
		fmt.Println(res)
	default:
		fmt.Println(errors.New("server error"))
	}
}

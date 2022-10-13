package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Write correct!")
		return
	}
	num1, err := strconv.Atoi(args[0])
	if err != nil {
		log.Printf("%v \nInter only number", err)
		return
	}
	oper := args[1]
	num2, err := strconv.Atoi(args[2])
	if err != nil {
		log.Println(err)
		return
	}
	switch oper {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("No divide by 0")
			return
		}
		fmt.Println(num1 / num2)
	case "%":
		if num2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		fmt.Println(num1 % num2)
	default:
		fmt.Println("Invalid operation")
	}

}

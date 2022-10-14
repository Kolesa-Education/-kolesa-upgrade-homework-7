package main

import (
	"calc/internal"
	"fmt"
	"log"
	"os"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	_, err := fmt.Fscan(os.Stdin, &num1)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Fscan(os.Stdin, &num2)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	fmt.Print("Enter the operator: ")
	_, err = fmt.Fscan(os.Stdin, &operator)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	res, err := internal.Calculate(num1, num2, operator)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", res)
	}
}

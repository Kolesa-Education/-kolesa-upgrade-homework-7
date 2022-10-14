package cli

import (
	calc "app/src/Calculator"
	"fmt"
)

func Run() {
	fmt.Print("For begin enter any value from the keyboard, to complete - 'exit': ")
	var cliOperator string

	for fmt.Scanln(&cliOperator); cliOperator != "exit"; fmt.Scanln(&cliOperator) {
		if cliOperator == "exit" {
			return
		}
		var operand1 float32
		var operand2 float32
		var operator string
		var res float32

		fmt.Print("Enter first operand: ")
		_, err := fmt.Scanln(&operand1)
		if err != nil {
			fmt.Println("invalid argument value entered, must be float")
			continue
		}
		fmt.Print("Enter first operand: ")
		_, err = fmt.Scanln(&operand2)
		if err != nil {
			fmt.Println("invalid argument value entered, must be float")
			continue
		}
		fmt.Print("Enter operator: ")
		fmt.Scanln(&operator)

		res, err = calc.Calculate(operand1, operand2, operator)
		if err != nil {
			fmt.Print(err)
			continue
		}
		fmt.Println(res)

		fmt.Print("To continue enter any value from the keyboard, to complete - 'exit': ")
		fmt.Scanln(cliOperator)
	}
}

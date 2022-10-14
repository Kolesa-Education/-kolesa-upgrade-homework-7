package main
import (
    "fmt"
    "math"
    "os"
)

func main() {
    var operand1 float64
    var operand2 float64
    var operator byte
    var res float64
    
    fmt.Printf("Enter first number: ")
    _, err := fmt.Scanf("%v", &operand1)
	if err != nil {
		fmt.Println("Incorrect first operand, please enter number")
		os.Exit(0)
	}
	
	fmt.Printf("Enter second number: ")
    _, err = fmt.Scanf("%v", &operand2)
	if err != nil {
		fmt.Println("Incorrect second operand, please enter number")
		os.Exit(0)
	}

    fmt.Printf("Enter operator: ")
    fmt.Scanf("%c", &operator)
    if (operator != '+' && operator != '-' && operator != '*' && operator != '/' && operator != '^' && operator != '>' && operator != '<' && operator != '=' && operator != '%') {
        fmt.Printf("Incorrect operator")
        os.Exit(0)
    }
    
    switch operator {
        case '+':
            res = operand1 + operand2
        case '-':
            res = operand1 - operand2
        case '*':
            res = operand1 * operand2
        case '/':
            res = operand1 / operand2
        case '^':
            res = math.Pow(operand1,operand2)
        case '<':
            if operand1 < operand2 {
                res = 1
            } else {
                res = 0
            }
        case '>':
            if operand1 > operand2 {
                res = 1
            } else {
                res = 0
            }
        case '=':
            if operand1 == operand2 {
                res = 1
            } else {
                res = 0
            }
        case '%':
            res = operand1 * operand2 / 100
    }
    
    fmt.Printf("Result: %v \n", res)
}
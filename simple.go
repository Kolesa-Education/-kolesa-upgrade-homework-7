package main
import (
	"fmt"
	"math"
)

func compare(a, b float64) string {
	if a > b {
        return "True";
    } 
	if a == b {
        return "They are equal";
    }
	return "False";
}


func main() {
	var (
		first float64
		operator string
		second float64
	)
	result := 0.0
	answer := ""
	
	fmt.Println("Enter first number: ")
	fmt.Scanln(&first)
	fmt.Println("Enter operator: ")
	fmt.Scanln(&operator)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&second)
	fmt.Println(first)

	switch operator {
	case "+":
		result = first + second
	case "-":
		result = first - second
	case "*":
		result = first * second
	case "/":
		if second == 0 {
			answer = "Division to 0 is not supported"
			break
		}		
		result = first / second
	case ">":
		answer = compare(first, second)
	case "<":
		answer = compare(second, first)
	case "^":
		result = math.Pow(first, second)
	case "%":
		result = first * second / 100;
	}
	if answer == "" {
		fmt.Println(result)
	} else {
		fmt.Println(answer)
	}
}
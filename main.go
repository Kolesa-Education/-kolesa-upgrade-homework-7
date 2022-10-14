package main
import (
	"fmt" 
	"math"
)

func main() {
	var num1 int
	var num2 int
	var op string
	fmt.Print("Enter 1st number: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("ERROR! Not a number!")
		return
	}
	fmt.Print("Enter 2nd number: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("ERROR! Not a number!")
		return
	}
	fmt.Print("Enter an operation: ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		fmt.Println( num1 + num2 )
	case "-":
		fmt.Println( num1 - num2 )
	case "/":
		if num2 == 0 {
			fmt.Println("Cannot divide to zero!");
			break
		}
		fmt.Println( float64(num1) / float64(num2) )
	case "*":
		fmt.Println( num1 * num2 )
	case "^":
		fmt.Println( math.Pow(float64(num1), float64(num2)) )
	case "%":
		fmt.Println( num1 % num2 )
	case "<":
		if num1 < num2 {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	case ">":
		if num1 > num2 {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	case "=":
		if num1 == num2 {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	default:
		fmt.Println("No such operation!")
	}
}
package main

import (
	"fmt"
	"log"
)

func operation(num1, num2, res float64, operator string) string {
	return fmt.Sprintf("%f, %s %f = %f", num1, operator, num2, res)
}

func main() {
	var num1, num2 float64
	var operator string

	_, err := fmt.Scanln(&num1)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = fmt.Scanln(&operator)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = fmt.Scanln(&num2)
	if err != nil {
		log.Println(err)
		return
	}

	switch operator {
	case "+":
		fmt.Printf("%v\n", operation(num1, num2, num1+num2, operator))
		break
	case "-":
		fmt.Printf("%v\n", operation(num1, num2, num1-num2, operator))
		break
	case "*":
		fmt.Printf("%v\n", operation(num1, num2, num1*num2, operator))
		break
	case "/":
		fmt.Printf("%v\n", operation(num1, num2, num1/num2, operator))
		break
	default:
		fmt.Println("Не совместивый оператор")
		break
	}

}

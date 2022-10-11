package main

import (
	"bufio"
	"fmt"
	"go/token"
	"go/types"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	var input string
	var str_number1, str_number2 string

	fmt.Println("Вы хотите вычислить простое выражение из двух операндов (1) или сложное (2)?")
	fmt.Println("'Задание часть 1' - (1) простое выраежение : 1 * 2.3")
	fmt.Println("'Задание часть 2' - (2) сложное выраежение : 1+3*(1+2/1-(2.5-1))")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	fmt.Println()
	if input == "1" {
		var operator string
		var number1, number2 float64

		fmt.Println("Можно использовать следующие операторы: +,-,/,*,^,>,<,=")

		scanner.Scan()
		res := strings.Split(scanner.Text(), " ")

		if len(res) != 3 {
			println(scanner.Text())
			println("неправильное выражение")
			return
		}

		str_number1 = res[0]
		operator = res[1]
		str_number2 = res[2]

		number1, err := strconv.ParseFloat(str_number1, 64)

		if err != nil {
			fmt.Printf("'%s' - это не число", str_number1)
			return
		}

		number2, err = strconv.ParseFloat(str_number2, 64)

		if err != nil {
			fmt.Printf("'%s' - это не число", str_number2)
			return
		}

		output := 0.0
		error := ""
		result := ""
		switch operator {
		case "+":
			output = number1 + number2
		case "-":
			output = number1 - number2
		case "*":
			output = number1 * number2
		case ">":
			if number1 > number2 {
				result = "true"
			} else {
				result = "false"
			}
		case "<":
			if number1 < number2 {
				result = "true"
			} else {
				result = "false"
			}
		case "=":
			if number1 == number2 {
				result = "true"
			} else {
				result = "false"
			}
		case "/":
			if number2 == 0.0 {
				error = "делить на 0 нельзя"
			} else {
				output = number1 / number2
			}
		case "^":
			if (number1 == 0.0) && (number2 == 0.0) {
				error = "неопределённость"
			} else {
				output = math.Pow(number1, number2)
			}
		default:
			error = "неправильная операция"
		}
		if error != "" {
			fmt.Println(error)
		} else if result != "" {
			fmt.Printf("%f %s %f is %s", number1, operator, number2, result)
		} else {
			fmt.Printf("%f %s %f = %f", number1, operator, number2, output)
		}
	} else if input == "2" {

		scanner.Scan()
		fs := token.NewFileSet()
		tv, err := types.Eval(fs, nil, token.NoPos, scanner.Text())
		if err != nil {
			println("ты делаешь то чего делать не надо! подумай над этим")
		}
		println("Ответ:", tv.Value.String())
	} else {
		println("Вы ввели что-то иное нежели 1 или 2. Такое моя не понимай!")
	}

}

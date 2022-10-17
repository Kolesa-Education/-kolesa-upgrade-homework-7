package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

func validate(elements[] string) (bool){
	if len(elements) != 3 {
		fmt.Println("Выражение не полное или содержит слишком много перепенных")
		return false
	}
	if _, err := strconv.ParseFloat(elements[0], 64); err != nil {
		fmt.Printf("%q Первый элемент не является числом.\n", elements[0])
		return false
	}

	if val, err := strconv.ParseFloat(elements[2], 64); err != nil {
		fmt.Printf("%q Второй элемент не является числом.\n", elements[2])

		return false
	}else if val == 0 && elements[1] == "/" {
		fmt.Println("Второй элемент при делений не может быть равен 0")
	}


	return true
}

func simple_calc(elements[] string){
	if validate(elements){
		res := 0.0
		value1, _ := strconv.ParseFloat(elements[0], 64)
		value2, _ := strconv.ParseFloat(elements[2], 64)
		operator := elements[1]

		switch operator{
		case "+":
			res = value1 + value2
		case "-":
			res = value1 - value2
		case "*":
			res = value1 * value2
		case "/":
			res = value1 / value2
		case "^":
			res = math.Pow(value1,value2)
		default:
			fmt.Println("Данный оператор не поддерживается")
		}
		fmt.Println(res)
	}else{
		fmt.Println("Ваше выражение не верное проверьте его еще раз")
	}
	
}


func main() {
	var expression string
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expression = scanner.Text()
	fmt.Print(expression, " = ")
	elements := strings.Split(expression, " ")

	simple_calc(elements)
}	

package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseArguments(chars []string) (float64, float64, string, error) {
	var num1, num2 float64
	var nums []float64
	var operator string

	for _, char := range chars {
		if char == " " {
		} else {
			num, err := strconv.ParseFloat(char, 64)
			if err != nil {
				operator += char
			} else {
				nums = append(nums, num)
			}
		}
	}
	if len(nums) < 2 {
		return 0.0, 0.0, "", errors.New("error: не хватает аргументов!")
	}
	num1 = nums[0]
	num2 = nums[1]
	if operator == "" {
		return 0.0, 0.0, "", errors.New("error: не хватает аргументов!")
	}

	return num1, num2, operator, nil
}

func calculating(expression string) (float64, error) {
	result := 0.0
	chars := strings.Split(expression, "")
	num1, num2, operator, err := parseArguments(chars)
	if err != nil {
		return 0.0, err
	}
	switch operator {
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0.0 {
			return 0.0, errors.New("error: деление на ноль!")
		}
		result = num1 / num2
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "^":
		result = math.Pow(num1, num2)
	case "%":
		result = num1 * (num2 / 100)
	case "==":
		if num1 >= num2 {
			result = num1
		} else {
			result = num2
		}
	default:
		return 0.0, errors.New(operator + " не является оператором!")
	}
	return result, nil
}

func main() {
	scanner := bufio.NewScanner((os.Stdin))
	input := make([]string, 0)
	fmt.Print("Введите выражение : \n")
	for {
		scanner.Scan()
		txt := scanner.Text()
		input = append(input, txt)
		res, err := calculating(input[len(input)-1])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
		fmt.Print("Введите выражение : \n")
	}

}

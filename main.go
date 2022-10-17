package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseArguments(chars []string) (float64, float64, string, error) {
	var num1, num2 float64
	var nums []float64
	var operator string
	var fullNum string
	i := 0
	for _, char := range chars {
		c, _ := regexp.MatchString("[0-9]", char)
		if c || char == "." {
			fullNum += char
			if i == len(chars)-1 {
				num, err := strconv.ParseFloat(fullNum, 64)
				if err != nil {
					return 0.0, 0.0, "", errors.New("неправильный ввод!")
				} else {
					nums = append(nums, num)
					fullNum = ""
				}
			}
		} else {
			if fullNum != "" {
				num, err := strconv.ParseFloat(fullNum, 64)
				if err != nil {
					return 0.0, 0.0, "", errors.New("неправильный ввод!")
				} else {
					nums = append(nums, num)
					fullNum = ""
				}
			}

			if char != " " {
				operator = char
			}
		}
		i++
	}

	if len(nums) < 2 {
		return 0.0, 0.0, "", errors.New("error: не хватает аргументов!")
	}
	if len(nums) > 2 {
		return 0.0, 0.0, "", errors.New("error: too much аргументов!")
	}
	num1 = nums[0]
	num2 = nums[1]
	if operator == "" {
		return 0.0, 0.0, "", errors.New("error: не хватает оператора!")
	}
	if len(operator) > 1 {
		return 0.0, 0.0, "", errors.New("error: слишком много операторов!")
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
		if len(operator) > 1 {
			return 0.0, errors.New("Слишком много аргументов!")
		} else {
			return 0.0, errors.New(operator + " не является оператором!")
		}
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

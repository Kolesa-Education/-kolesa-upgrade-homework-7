package main

import (
	"fmt"
	"strconv"
)

func calc(text string) string {
	num1 := 0
	num2 := 0
	turn := false
	var c rune
	for _, val := range text {
		if val >= '0' && val <= '9' {
			num, _ := strconv.Atoi(string(val))
			if turn == false {
				num1 = num1*10 + num
			} else {
				num2 = num2*10 + num
			}
		} else {
			if turn == false {
				turn = true
			}
			c = val
		}
	}
	res := 0.0
	if c == '+' {
		num1 += num2
	} else if c == '-' {
		num1 -= num2
	} else if c == '*' {
		num1 *= num2
	} else {
		if c == '/' {
			if num2 == 0 {
				return "Error: divide by zero!"
			}
			res = float64(num1) / float64(num2)
			return fmt.Sprintf("%f", res)
		} else {
			return "Error: wrong input!"
		}
	}
	return strconv.Itoa(num1)
}

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(calc(text))
}

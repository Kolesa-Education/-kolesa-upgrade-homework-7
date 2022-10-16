package internal

import (
	"strconv"
	"strings"
)

var opMap = map[string]int{
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
	"(": -1,
	")": -1,
}

func CalculatePostfix() float64 {
	st_calc := []float64{}
	for _, token := range tokens_postfix {
		if _, found := opMap[token]; found {
			a := st_calc[len(st_calc)-1]
			b := st_calc[len(st_calc)-2]
			switch token[0] {
			case '+':
				b += a
				break
			case '-':
				b -= a
				break
			case '*':
				b *= a
				break
			case '/':
				b /= a
			}
			st_calc = st_calc[:len(st_calc)-2]
			st_calc = append(st_calc, b)
		} else {
			if val, err := strconv.ParseFloat(token, 64); err == nil {
				st_calc = append(st_calc, val)
			}
		}
	}
	return st_calc[0]
}

func Calculate(expression string) (res float64) {
	if !strings.ContainsAny(expression, "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz") {
		Tokenize(expression)
		InfixToPostfix()
		res = CalculatePostfix()
	}
	return

}

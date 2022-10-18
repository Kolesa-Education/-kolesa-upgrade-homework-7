package internal

import (
	"strconv"
	"strings"
)

func CalculatePostfix(postfix []string, opMap map[string]int) float64 {
	st_calc := []float64{}
	for _, token := range postfix {
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
	blackList := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	if strings.ContainsAny(expression, blackList) {
		return
	}
	var opMap = map[string]int{
		"*": 2,
		"/": 2,
		"+": 1,
		"-": 1,
		"(": -1,
		")": -1,
	}
	infix := Tokenize(expression, opMap)
	postfix := InfixToPostfix(infix, opMap)
	res = CalculatePostfix(postfix, opMap)
	return

}

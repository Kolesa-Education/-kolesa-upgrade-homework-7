package main

import (
	"errors"
	"math"
)

func Calc(a, b float64, operator string) (interface{}, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("undefined")
		} else {
			return a / b, nil
		}
	case "*":
		return a * b, nil
	case "%":
		return math.Mod(a, b), nil
	case ">":
		if a > b {
			return true, nil
		}
		return false, nil
	case "<":
		if a < b {
			return true, nil
		}
		return false, nil
	case "=":
		if a == b {
			return true, nil
		}
		return false, nil
	default:
		return 0, errors.New("invalid operator")
	}
}

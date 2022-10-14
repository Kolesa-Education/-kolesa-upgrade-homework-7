package calc

import (
	"errors"
	"math"
)

func Calculate(operand1 float32, operand2 float32, operator string) (float32, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			panic("cannot be divided by zero")
		}
		return operand1 / operand2, nil
	case "^":
		return float32(math.Pow(float64(operand1), float64(operand2))), nil
	case "?":
		if operand1 >= operand2 {
			return operand1, nil
		}
		return operand2, nil
	case "%":
		return operand1 * operand2 / 100, nil
	default:
		return 0, errors.New("invalid operator")
	}
}

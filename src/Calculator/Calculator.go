package calc

import (
	"errors"
	"github.com/maja42/goval"
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
			return 0, errors.New("cannot be divided by zero")
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

func CalculateExpr(expr string) (interface{}, error) {
	eval := goval.NewEvaluator()
	res, err := eval.Evaluate(expr, nil, nil)
	if err != nil {
		return nil, errors.New("incorrect expression")
	}
	return res, nil
}

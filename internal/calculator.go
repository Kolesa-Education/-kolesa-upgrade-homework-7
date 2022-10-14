package internal

import (
	"errors"
	"math"
)

func Calculate(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return add(a, b), nil
	case "-":
		return substract(a, b), nil
	case "*":
		return multiply(a, b), nil
	case "/":
		return divide(a, b)
	case "mod":
		return modulo(int(a), int(b))
	case "^":
		return math.Pow(a, b), nil
	case "%":
		return percent(a, b)
	default:
		return 0, errors.New("Not valid operation")
	}
}

func add(a, b float64) float64 {
	return a + b
}

func substract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, errors.New("No division by 0")
	}

	return a / b, nil
}

func modulo(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("No modulo by 0")
	}

	return float64(a % b), nil
}

func percent(a, b float64) (float64, error) {
	if b < 0 {
		return 0, errors.New("Percentage can't be negative")
	}

	return (a * b) / 100.0, nil
}

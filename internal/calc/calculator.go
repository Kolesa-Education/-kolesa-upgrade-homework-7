package calc

import (
	"fmt"
	"kolesa-upgrade-homework-7/internal/parser"
	"log"
	"math"
	"strconv"
)

func Calculate(tokens []parser.Token) float64 {

	var stack []float64

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type == parser.NUM {
			num, err := strconv.ParseFloat(tokens[i].Value, 64)
			if err != nil {
				log.Fatalln("non-number argument", err)
			}
			stack = append(stack, num)
		} else {
			var leftOp, rightOp float64
			leftOp, stack = PopFloatStack(stack)
			rightOp, stack = PopFloatStack(stack)

			switch tokens[i].Value {
			case "+":
				res := rightOp + leftOp
				stack = append(stack, res)
			case "-":
				res := rightOp - leftOp
				stack = append(stack, res)
			case "/":
				res := rightOp / leftOp
				stack = append(stack, res)
			case "*":
				res := rightOp * leftOp
				stack = append(stack, res)
			case "^":
				res := math.Pow(rightOp, leftOp)
				stack = append(stack, res)
			}
		}
	}
	fmt.Println(len(stack))
	res, _ := PopFloatStack(stack)

	return res
}

func PopFloatStack(stack []float64) (float64, []float64) {
	var element float64
	if len(stack) > 0 {
		element = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return element, stack
}

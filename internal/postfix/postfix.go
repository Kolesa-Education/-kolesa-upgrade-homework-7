package postfix

import (
	"fmt"
	"kolesa-upgrade-homework-7/internal/parser"
)

var OperPrecedence = map[string]int{
	"+": 20,
	"-": 20,
	"*": 30,
	"/": 30,
	"^": 40,
	"(": 10,
}

func PostFix(tokens []parser.Token) []parser.Token {
	var resQueue []parser.Token
	var opStack []parser.Token

	topStackPrecedence := 0

	for i := 0; i < len(tokens); i++ {
		fmt.Println(tokens[i], topStackPrecedence)
		if tokens[i].Type == parser.NUM {
			resQueue = enqueue(resQueue, tokens[i])
		} else if tokens[i].Type == parser.OP && len(opStack) == 0 {
			opStack = Push(opStack, tokens[i])
			topStackPrecedence = OperPrecedence[opStack[len(opStack)-1].Value]
		} else if tokens[i].Type == parser.OP && OperPrecedence[tokens[i].Value] > topStackPrecedence {
			opStack = Push(opStack, tokens[i])
			topStackPrecedence = OperPrecedence[opStack[len(opStack)-1].Value]
		} else if tokens[i].Type == parser.OP && tokens[i].Value == "(" {
			opStack = Push(opStack, tokens[i])
			topStackPrecedence = OperPrecedence[opStack[len(opStack)-1].Value]

		} else if tokens[i].Type == parser.OP && tokens[i].Value == ")" {
			var token parser.Token

			token, opStack = Pop(opStack)

			for token.Value != "(" && len(opStack) != 0 {
				resQueue = enqueue(resQueue, token)
				token, opStack = Pop(opStack)
			}

			if len(opStack) > 0 {
				topStackPrecedence = OperPrecedence[opStack[len(opStack)-1].Value]
			}

		} else {
			for OperPrecedence[tokens[i].Value] <= topStackPrecedence {
				token, newStack := Pop(opStack)
				opStack = newStack

				resQueue = enqueue(resQueue, token)

				if len(opStack) > 0 {
					topStackPrecedence = OperPrecedence[opStack[len(opStack)-1].Value]
				} else {
					topStackPrecedence = 0
				}
			}

			opStack = Push(opStack, tokens[i])
			topStackPrecedence = OperPrecedence[opStack[len(opStack)-1].Value]
		}

	}

	for len(opStack) > 0 {
		var token parser.Token
		token, opStack = Pop(opStack)

		resQueue = append(resQueue, token)
	}

	return resQueue
}

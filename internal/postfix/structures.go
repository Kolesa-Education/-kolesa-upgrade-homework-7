package postfix

import "kolesa-upgrade-homework-7/internal/parser"

func enqueue(queue []parser.Token, element parser.Token) []parser.Token {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []parser.Token) (parser.Token, []parser.Token) {
	return queue[0], queue[1:]
}

func Push(stack []parser.Token, element parser.Token) []parser.Token {
	stack = append(stack, element)
	return stack
}

func Pop(stack []parser.Token) (parser.Token, []parser.Token) {
	var element parser.Token
	if len(stack) > 0 {
		element = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	return element, stack
}

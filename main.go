package main

import (
	"bufio"
	"errors"
	"fmt"
	"kolesa-calc/stack_kolesa"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var operators = map[string]struct {
	precedency int
	power      bool
}{
	"^": {4, true},
	"*": {3, false},
	"/": {3, false},
	"+": {2, false},
	"-": {2, false},
}

func isOperator(ch rune) bool {
	if ch == '+' || ch == '-' || ch == '/' || ch == '*' || ch == '^' {
		return true
	}
	return false
}

func isBracket(ch rune) bool {
	if ch >= '(' && ch <= ')' {
		return true
	}
	return false
}
func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func isValid(str string) bool {
	for _, sym := range str {
		if !isDigit(sym) && !isOperator(sym) &&
			!isBracket(sym) && sym != ' ' {
			fmt.Println(sym)
			return false
		}
	}
	return true
}
func spaces(str string) (string, error) {
	if !isValid(str) {
		return "", errors.New("non-valid string")
	}
	str = strings.TrimSpace(str)
	res := string(str[0])
	prevNumber := false
	if isDigit(rune(str[0])) {
		prevNumber = true
	}
	for i := 1; i < len(str); i++ {
		if isDigit(rune(str[i])) && prevNumber || rune(str[i]) == ' ' {
			res += string(str[i])
			continue
		}
		if !isDigit(rune(str[i])) {
			prevNumber = false
		} else {
			prevNumber = true
		}
		res += " " + string(str[i])
	}
	return res, nil
}

func parseInfix(e string) (rpn string) {
	stack := stack_kolesa.Stack{}
	for _, tok := range strings.Fields(e) {
		switch tok {
		case "(":
			stack.Push(tok)
		case ")":
			var op string
			for {
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break
				}
				rpn += " " + op
			}
		default:
			if o1, isOp := operators[tok]; isOp {
				for len(stack) > 0 {
					op := stack[len(stack)-1]
					if o2, isOper := operators[op]; !isOper ||
						o1.precedency > o2.precedency ||
						o1.precedency == o2.precedency &&
							o1.power {
						break
					}
					_, err := stack.Pop()
					if err != nil {
						log.Println(err)
						return
					}
					rpn += " " + op
				}
				stack.Push(tok)
			} else {
				if rpn > "" {
					rpn += " "
				}
				rpn += tok
			}
		}
	}
	for len(stack) > 0 {
		rpn += " " + stack[len(stack)-1]
		_, err := stack.Pop()
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

func rpn(tks []string) {
	var (
		nums []float64
		x    float64
	)
	pop := func() float64 {
		n := len(nums) - 1
		t := nums[n]
		nums = nums[:n]
		return t
	}
	defer func() {
		if recover() == nil && len(nums) == 1 {
			fmt.Println(x)
		} else {
			fmt.Println("error")
		}
	}()
	for _, tk := range tks {
		switch tk {
		case "+":
			x = pop() + pop()
		case "*":
			x = pop() * pop()
		case "-":
			x = pop()
			x = pop() - x
		case "/":
			x = pop()
			x = pop() / x
		case "^":
			y := pop()
			x = pop()

			for i := 0.0; i < y-1; i++ {
				x *= x
			}

		default:
			var e error
			x, e = strconv.ParseFloat(tk, 64)
			if e != nil {
				panic(0)
			}
		}
		nums = append(nums, x)
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	input = strings.TrimSpace(input)
	str, err := spaces(input)
	if err != nil {
		log.Fatal(err)
	}
	rpn(strings.Fields(parseInfix(str)))
}

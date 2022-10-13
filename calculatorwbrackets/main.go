package main

import (
	"fmt"
)

type Dek struct {
	operands  []int
	operators []byte
}

func NewDek() *Dek {
	return &Dek{
		operands:  make([]int, 0),
		operators: make([]byte, 0),
	}
}

func (d *Dek) PushOperand(val int) {
	d.operands = append(d.operands, val)
}

func (d *Dek) PopOperand() int {
	val := d.operands[len(d.operands)-1]
	d.operands = d.operands[:len(d.operands)-1]

	return val
}

func (d *Dek) PushOperator(ch byte) {
	d.operators = append(d.operators, ch)
}

func (d *Dek) PopOperator() byte {
	val := d.operators[len(d.operators)-1]
	d.operators = d.operators[:len(d.operators)-1]

	return val
}

func (d *Dek) PeekOperator() byte {
	return d.operators[len(d.operators)-1]
}

func (d *Dek) isOperatorEmpty() bool {
	return len(d.operators) == 0
}

func (d *Dek) operate() int {
	a, b, c := d.PopOperand(), d.PopOperand(), d.PopOperator()

	switch c {
	case '+':
		return a + b
	case '-':
		return b - a
	case '*':
		return a * b
	case '/':
		return b / a
	default:
		return 0
	}
}

var priority = map[byte]int{
	'(': -1,
	'+': 0,
	'-': 0,
	'*': 1,
	'/': 1,
}

func calculate(s string) int {
	nd, n := NewDek(), len(s)
	for i := 0; i < n; i++ {
		x := s[i]
		if isDigit(x) {
			val := int(x - '0')
			for i+1 < n && isDigit(s[i+1]) {
				val = val*10 + int(s[i+1]-'0')
				i++
			}
			nd.PushOperand(val)
		} else if x == ' ' {
			continue
		} else if x == '(' {
			nd.PushOperator(x)
		} else if x == ')' {
			for nd.PeekOperator() != '(' {
				nd.PushOperand(nd.operate())
			}
			nd.PopOperator()
		} else {
			for !nd.isOperatorEmpty() && comparePriority(x, nd.PeekOperator()) <= 0 {
				nd.PushOperand(nd.operate())
			}
			nd.PushOperator(x)
		}
	}
	for !nd.isOperatorEmpty() {
		nd.PushOperand(nd.operate())
	}
	return nd.PopOperand()
}

func comparePriority(a, b byte) int {
	return priority[a] - priority[b]
}

func isDigit(x byte) bool {
	if x >= '0' && x <= '9' {
		return true
	}
	return false
}

func main() {
	var input string
	fmt.Println("Введите числовое выражение со скобками:")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Ответ: ", calculate(input))
}

//   1+3*(1+2/1-(2-1))

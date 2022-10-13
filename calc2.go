package main

import (
	"errors"
	"fmt"
)

type tokenType int

const (
	leftParent  tokenType = iota // (
	rightParent                  // )
	minus                        // -
	plus                         // +
	slash                        // "/"
	star                         // '*'
	number                       // 123, 9, 10
	eof                          // end token type
)

type token struct {
	TkType  tokenType
	Lexeme  string
	Literal interface{}
}

type scanner struct {
	exp     string
	runes   []rune   // the chars in the expression
	tokens  []*token // the tokens collected from expression
	start   int      // scan start from
	current int      // current scan index
}

func newScanner(exp string) *scanner {
	return &scanner{
		exp:    exp,
		runes:  []rune(exp),
		tokens: make([]*token, 0, 10),
	}
}

var ErrDivideByZero = errors.New("divide by zero")

func Divide(exp string) (string, error) {
	if exp == slash {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {
	fmt.Println("Enter expression you want to calculate")
	result, err := Divide(a, b)
	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("divide by zero error")
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)
}

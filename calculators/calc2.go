package calculators

import "fmt"

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

type DivZero struct{}

func (myerr *DivZero) Error() string {
	return "Cannot divide by 0!"
}

func Calc2(text string) {

	fmt.Println("Enter expression you want to calculate")
}

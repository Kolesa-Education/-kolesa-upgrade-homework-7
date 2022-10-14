package parser

import (
	"regexp"
	"unicode"
)

func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func IsWhiteSpace(r rune) bool {
	return unicode.IsSpace(r)
}

func IsOperand(r rune) bool {
	regOp := regexp.MustCompile(`\(|\)|\-|\+|\/|\*|\^|\%`)

	return regOp.MatchString(string(r))
}

func ValidTokens(tokens []Token) bool {
	countScope := 0
	prevOp := ""

	for _, token := range tokens {
		if token.Type == NUM {
			prevOp = ""
			continue
		}
		if token.Value == "(" {
			countScope++
		} else if token.Value == ")" {
			countScope--
		}

		if token.Value == "+" || token.Value == "-" || token.Value == "*" || token.Value == "/" ||
			token.Value == "^" || token.Value == ")" {
			if prevOp != ")" && prevOp != "" {
				return false
			}
		}
		prevOp = token.Value
	}

	if countScope != 0 {
		return false
	}
	return true
}
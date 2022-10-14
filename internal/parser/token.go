package parser

type tokenType int

const (
	NUM = iota //number
	OP         //operand
)

type Token struct {
	Type  tokenType
	Value string
}

func ConvertToTokens(s string) []Token {
	var tokens []Token

	runes := []rune(s)
	start := 0

	for i := 0; i < len(runes); i++ {
		if IsDigit(runes[i]) || IsWhiteSpace(runes[i]) {
			continue
		}
		if IsOperand(runes[i]) {
			if i > 0 && start != i {
				tokenNum := Token{Type: NUM, Value: string(runes[start:i])}
				tokens = append(tokens, tokenNum)
			}
			tokenOp := Token{Type: OP, Value: string(runes[i])}
			tokens = append(tokens, tokenOp)

			start = i + 1

		}
	}

	if start < len(runes) {
		tokens = append(tokens, Token{Type: NUM, Value: string(runes[start:])})
	}

	return tokens
}

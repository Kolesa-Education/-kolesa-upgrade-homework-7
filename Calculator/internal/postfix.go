package internal

var tokens_postfix []string

func InfixToPostfix() {
	st := []string{"("}
	tokens_infix = append(tokens_infix, ")")
	for _, token := range tokens_infix {
		if _, found := opMap[token]; !found { // operand
			tokens_postfix = append(tokens_postfix, token)
		} else if token == "(" {
			st = append(st, token)
		} else if token == ")" {
			for st[len(st)-1] != "(" {
				tokens_postfix = append(tokens_postfix, st[len(st)-1])
				st = st[:len(st)-1]
			}
			st = st[:len(st)-1]
		} else {
			for len(st) > 0 && opMap[token] <= opMap[st[len(st)-1]] {
				tokens_postfix = append(tokens_postfix, st[len(st)-1])
				st = st[:len(st)-1]
			}
			st = append(st, token)
		}
	}

}

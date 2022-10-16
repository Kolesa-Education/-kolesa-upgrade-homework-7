package internal

var tokens_infix []string

func Tokenize(expression string) {
	temp := ""
	for _, ch := range expression {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			panic("Invalid expression")
		}
		if ch == ' ' || ch == '\n' {
			continue
		} else if _, found := opMap[string(ch)]; found { // v =value found == true
			if temp != "" {
				tokens_infix = append(tokens_infix, temp)
			}
			temp = ""
			tokens_infix = append(tokens_infix, string(ch))
		} else {
			temp += string(ch)
		}
	}
	if temp != "" {
		tokens_infix = append(tokens_infix, temp)
	}
}

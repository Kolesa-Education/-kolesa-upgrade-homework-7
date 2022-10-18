package internal

import "fmt"

func Tokenize(expression string, opMap map[string]int) []string {
	var infix []string
	token := ""
	for _, ch := range expression {
		if ch == ' ' || ch == '\n' {
			continue
		} else if _, found := opMap[string(ch)]; found { // v =value found == true
			if token != "" {
				infix = append(infix, token)
			}
			token = ""
			infix = append(infix, string(ch))
		} else {
			token += string(ch)
		}
	}
	if token != "" {
		infix = append(infix, token)
	}
	fmt.Println(infix)
	return infix
}

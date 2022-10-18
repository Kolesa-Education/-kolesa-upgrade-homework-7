package internal

import "fmt"

func InfixToPostfix(infix []string, opMap map[string]int) []string {
	var postfix []string
	st := []string{"("}
	infix = append(infix, ")")
	for _, token := range infix {
		if _, found := opMap[token]; !found { // operand
			postfix = append(postfix, token)
		} else if token == "(" {
			st = append(st, token)
		} else if token == ")" {
			for st[len(st)-1] != "(" {
				postfix = append(postfix, st[len(st)-1])
				st = st[:len(st)-1]
			}
			st = st[:len(st)-1]
		} else {
			for len(st) > 0 && opMap[token] <= opMap[st[len(st)-1]] {
				postfix = append(postfix, st[len(st)-1])
				st = st[:len(st)-1]
			}
			st = append(st, token)
		}
	}
	fmt.Println(postfix)
	return postfix
}

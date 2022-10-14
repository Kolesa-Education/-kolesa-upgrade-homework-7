package main

import (
	"bufio"
	"fmt"
	"kolesa-upgrade-homework-7/internal/calc"
	"kolesa-upgrade-homework-7/internal/parser"
	"kolesa-upgrade-homework-7/internal/postfix"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Write your expression: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	text = strings.TrimSpace(text)
	tokens := parser.ConvertToTokens(strings.ReplaceAll(text, " ", ""))

	fmt.Println(tokens)

	//if !parser.ValidTokens(tokens) {
	//	log.Fatalln("some lexical error in given expression")
	//}
	fixedTokens := postfix.PostFix(tokens)

	fmt.Println(fixedTokens)

	res := calc.Calculate(fixedTokens)

	fmt.Println(res)

}

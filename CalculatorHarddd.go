package main

import (
	"bufio"
	"go/token"
	"go/types"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileSet := token.NewFileSet()
	evalRes, err := types.Eval(fileSet, nil, 0, scanner.Text())
	if err != nil {
		println("Что-то пошло не так...")
	}
	println("Ответ:", evalRes.Value.String())
}

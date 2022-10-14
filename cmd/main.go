package main

import (
	"bufio"
	"fmt"
	"kolesa-upgrade-homework-7/internal/calc"
	"kolesa-upgrade-homework-7/internal/parser"
	"kolesa-upgrade-homework-7/internal/postfix"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Write your expression: ")
	text, _ := reader.ReadString('\n') //Чтение математического выражения

	text = strings.TrimSpace(text) // убираем начальные и концовые пробелы

	// Конвертируем в токены
	tokens := parser.ConvertToTokens(strings.ReplaceAll(text, " ", ""))

	//Проверяем токены: наличия не равных количеств скоб для открытия и закрытия,
	//лексические ошибки как два операнда (пример: ++, /+ etc.)
	if !parser.ValidTokens(tokens) {
		log.Fatalln("some lexical error in given expression")
	}

	//Получаем постфикс нотацию токенов
	fixedTokens := postfix.PostFix(tokens)
	//Вычислить выражение
	res := calc.Calculate(fixedTokens)

	fmt.Println(res)
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string = ""
	println("Waiting for input...")
	fmt.Scanln(&input)
	if input == "" {
		println("No value provided")
		return
	}
	var splittedInput []string = splitInput(input)
	checkInput(splittedInput)
	var unbracked []string = bracketSplitter(splittedInput)
	var result string = calcValues(unbracked)
	println(result)
}

// Разделить строку на массив знаков и чисел
func splitInput(input string) []string {
	input = strings.ReplaceAll(input, " ", "")
	separated := strings.Split(input, "")
	var splitted []string = nil
	for _, str := range separated {
		//Если не число- добавляем в слайс и продолжаем
		if _, err := strconv.Atoi(str); err != nil {
			splitted = append(splitted, str)
			continue
		}
		if len(splitted) != 0 {
			//Если предыдущий добавленный элемент- число, объединяем строки и продолжаем
			if _, e := strconv.Atoi(splitted[len(splitted)-1]); e == nil {
				splitted[len(splitted)-1] += str
				continue
			}
		}
		splitted = append(splitted, str)
	}
	return splitted
}

// Валидация ввода
func checkInput(input []string) {
	var prevEl string = ""
	var ok bool = true
	if _, err := strconv.Atoi(input[0]); err != nil {
		if input[0] != "(" {
			ok = false
		} else {
			prevEl = "("
		}
	} else {
		prevEl = "number"
	}
	if _, err := strconv.Atoi(input[len(input)-1]); err != nil {
		if input[len(input)-1] != ")" {
			ok = false
		}
	}
	if ok == false {
		println("Incorrect input")
		os.Exit(1)
	}
	for i := 1; i < len(input); i++ {
		switch {
		case input[i] == "(":
			if prevEl != "operator" && prevEl != "(" {
				ok = false
			}
			prevEl = "("
		case input[i] == ")":
			if prevEl != "number" && prevEl != ")" {
				ok = false
			}
			prevEl = ")"
		case input[i] == "+" || input[i] == "-" || input[i] == "*" || input[i] == "/":
			if prevEl != "number" && prevEl != ")" {
				ok = false
			}
			prevEl = "operator"
		default:
			if _, err := strconv.Atoi(input[i]); err == nil {
				if prevEl != "operator" && prevEl != "(" {
					ok = false
				}
				prevEl = "number"
			}
		}
		if ok == false {
			println("Incorrect input")
			os.Exit(1)
		}
	}
}

// Функция для раскрытия скобок
func bracketSplitter(input []string) []string {
	//Символы вне скобок и раскрытые скобки
	var unbracked []string = nil
	//Операции в скобках
	var bracked []string = nil
	//Количество открытых скобок
	var brCount int = 0

	for _, str := range input {
		//Если элемент- открывающая скобка
		if str == "(" {
			//Если уже были открывающие скобки
			if brCount > 0 {
				//Добавляем в массив с нераскрытыми скобками
				bracked = append(bracked, str)
			}
			//Если это первая открывающая скобка- увеличиваем счетчик и пропускаем элемент
			brCount++
			continue
			//Если элемент- закрывающая скобка
		} else if str == ")" {
			//Уменьшаем счетчик открытых скобок
			brCount--
			//Если все скобки закрылись
			if brCount == 0 {
				//Отправляем в рекурсию массив нераскрытых строк
				res := bracketSplitter(bracked)
				//Считаем результат открытых скобок и добавляем значение в слайс
				unbracked = append(unbracked, calcValues(res))
			}
		}
		//Если есть открытые скобки
		if brCount > 0 {
			//Добавляем элемент в слайс с нераскрытыми скобками
			bracked = append(bracked, str)
		} else {
			//Если нет открытых скобок и элемент- не ")", добавляем значение к слайсу с раскрытыми скобками
			if str != ")" {
				unbracked = append(unbracked, str)
			}
		}
	}
	return unbracked
}

// Произволит расчет с учетом приоритета операций
func calcValues(input []string) string {
	//Результат после выполнения приоритетных операций
	var prior []string = nil
	//Финальный результат
	var result []string = nil

	//Поиск и выполнение приоритетных операций
	for i := 0; i < len(input); i++ {
		if input[i] == "*" || input[i] == "/" {
			num1, _ := strconv.Atoi(input[i-1])
			num2, _ := strconv.Atoi(input[i+1])
			prior[len(prior)-1] = strconv.Itoa(simpleCalc(num1, num2, input[i]))
			i++
			continue
		}
		prior = append(prior, input[i])
	}

	//Выполнение неприоритетных операций
	for i := 0; i < len(prior); i++ {
		if prior[i] == "+" || prior[i] == "-" {
			num1, _ := strconv.Atoi(result[len(result)-1])
			num2, _ := strconv.Atoi(prior[i+1])
			result[len(result)-1] = strconv.Itoa(simpleCalc(num1, num2, prior[i]))
			i++
			continue
		}
		result = append(result, prior[i])
	}
	return result[0]
}

// Простые операции над двумя числами
func simpleCalc(a, b int, c string) int {
	var result int = 0
	switch c {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			println("Division by zero")
			os.Exit(1)
		}
		result = a / b
	default:
		println("Unknown operation:", a, c, b)
		os.Exit(1)
	}
	println(a, c, b, "=", result)
	return result
}

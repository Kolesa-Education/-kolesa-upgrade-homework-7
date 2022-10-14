package main

import (
	"fmt"
	"kolesa-upgrade-homework-7/service"
)

func main() {
	var tempNum1 string
	var oper string
	var tempNum2 string
	fmt.Println("Введите математическое выражение через пробел")
	fmt.Scan(&tempNum1, &oper, &tempNum2)
	service.Calc(tempNum1, oper, tempNum2)
}

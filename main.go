package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	x, o, y := 0, "", 0
	counter := 0
loop:
	for {
		counter++
		fmt.Println(counter)
		fmt.Println("введите задачу в одну строку")
		fmt.Println("пример:")
		fmt.Println("2 + 5")
		task := ""
		n, err := fmt.Scanf("%d %s %d", &x, &o, &y)
		if err != nil || n != 3 {
			continue
		}
		task = strconv.Itoa(x) + o + strconv.Itoa(y)
		re := regexp.MustCompile(`\d+\s*(\+|-|\*|\/|%)\s*\d+`)
		fmt.Printf("------------------\ntask: %s\n------------------\n", task)
		match := re.Match([]byte(task))
		if match {
			break
		}
	}
	fmt.Printf("%d %s %d = ", x, o, y)
	switch o {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		if x == 0 {
			fmt.Println("делить на ноль нельзя")
			goto loop
		}
		fmt.Println(x / y)
	case "%":
		fmt.Println(x % y)
	}
}

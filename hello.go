package main

import (
	"fmt"
	"math"
	"regexp"
	"runtime"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("-----")
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	var a, b int
	var c string
	match, _ := regexp.MatchString("^[a-zA-Z0-9]$", c)
	fmt.Scan(&a, &c, &b)
	if match {
		fmt.Println("не правильный оператор")
	} else {
		switch c {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		case "%":
			fmt.Println(a % b)
		case "**":
			fmt.Println(math.Pow(float64(a), float64(b)))
		default:
			fmt.Println("не правильный оператор")
		}

	}
	printStats(mem)
}

package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/m/v2/calculators"
)

func main() {
	var input string
	fmt.Println("Choose simple(1) or complex(2) calculator")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	if input == "1" {
		scanner.Scan()
		calculators.Calc1(scanner.Text())
	} else if input == "2" {
		scanner.Scan()
		calculators.Calc2(scanner.Text())
	} else {
		println("Try again! Choose 1 or 2")
	}
}

package main

import (
	"Calculator/internal"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter expression: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(internal.Calculate(text))

}

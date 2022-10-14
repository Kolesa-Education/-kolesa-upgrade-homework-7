package main

import (
	"fmt"
	"os"
	"strconv"
)
// test version. yesterday i will edit that, but  today i wrote simple-3-args calculator :(
func main() {
	if len(os.Args) != 4 {
		fmt.Println("incorrect structure")
		return
	}
	num1,err:=strconv.Atoi(os.Args[1])
	if err!=nil{
		fmt.Println("incorrect number")
		return
	}
	operation:=os.Args[2]
	num2,err:=strconv.Atoi(os.Args[3])
	if err!=nil{
		fmt.Println("incorrect number")
		return
	}
	result:=0
	switch operation{
	case "*":
		result=num1*num2
	case "/":
		result=num1/num2
	case "+":
		result=num1+num2
	case "-":
		result=num1-num2
	default:
		fmt.Println("incorrect operation")
		return

	}
	fmt.Println("result =",result)
	

}

package main

import (
	"fmt"
	"math"
	"strconv"
)

type Calculator struct{
	Num1, Num2 int
}

func (c Calculator) Plus(){
	fmt.Println(c.Num1+c.Num2)
}

func (c Calculator) Minus(){
	fmt.Println(c.Num1-c.Num2)
}

func (c Calculator) Multiply(){
	fmt.Println(c.Num1*c.Num2)
}

func (c Calculator) Divide(){
	defer func() {
		if r := recover(); r != nil{
			fmt.Println(r)
		}
	}()
	fmt.Println(c.Num1/c.Num2)
}

func (c Calculator) Degree(){
	fmt.Println(math.Pow(float64(c.Num1), float64(c.Num2)))
}

func (c Calculator) Comparison(){
	if c.Num1>c.Num2{
		fmt.Println("Num1 is greater than Num2")
	} else if c.Num1<c.Num2{
		fmt.Println("Num1 is less than Num2")
	} else{
		fmt.Println("Num1 equals Num2")
	}
}

func (c Calculator) Percent(){
	fmt.Println(float64(c.Num1)/float64(c.Num2)*100)
}

func main(){
	var a, b, c string
	fmt.Print("num1 = ")
	fmt.Scan(&a)
	fmt.Print("operator = ")
	fmt.Scan(&c)
	fmt.Print("num2 = ")
	fmt.Scan(&b)
	num1, err1 := strconv.Atoi(a)
	num2, err2 := strconv.Atoi(b)
	if err1 == nil && err2 == nil{ 
		numbers := Calculator{Num1 : num1, Num2 : num2,}
		switch c{
		case "+":
			numbers.Plus()
		case "-":
			numbers.Minus()
		case "*":
			numbers.Multiply()
		case "/":
			numbers.Divide()
		case "^":
			numbers.Degree()
		case "compare":
			numbers.Comparison()
		case "%":
			numbers.Percent()
		default:
			fmt.Println("Wrong operation")
		}
	} else{
		fmt.Println("Check the correctness of the entered numbers")
	}
}

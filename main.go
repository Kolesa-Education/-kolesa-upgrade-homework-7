package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseStrToSlice(my_str string) []string {
	my_str = strings.Replace(my_str, " ", "", -1)

	re, err := regexp.Compile(`[\d+]|[+]|[-]|[*]|[/]|[(]|[)]`)
	if err != nil {
		fmt.Println(err)
	}

	res := re.FindAllStringSubmatch(my_str, -1)
	ret := make([]string, len(res))
	for i, v := range res {
		ret[i] = v[0]
	}

	return ret
}

func removeIndex(my_slice []string, i int) []string {
	ret := make([]string, 0)
	ret = append(ret, my_slice[:i]...)
	return append(ret, my_slice[i+1:]...)
}

func removePart(my_slice []string, begin int, end int) []string {
	ret := make([]string, 0)
	ret = append(ret, my_slice[:begin]...)
	return append(ret, my_slice[end+1:]...)
}

func mult_div(my_slice []string) []string {
	var my_len = len(my_slice) - 1
	for i := 0; i < my_len; i++ {
		var x int
		var v string = my_slice[i]

		if v == "*" {
			a, _ := strconv.Atoi(my_slice[i-1])
			b, _ := strconv.Atoi(my_slice[i+1])
			x = a * b
		} else if v == "/" {
			a, _ := strconv.Atoi(my_slice[i-1])
			b, _ := strconv.Atoi(my_slice[i+1])
			x = a / b
		}

		if v == "*" || v == "/" {
			my_slice = removeIndex(my_slice, i-1)
			my_slice = removeIndex(my_slice, i-1)
			my_slice[i-1] = strconv.Itoa(x)
			return mult_div(my_slice)
		}
	}

	return my_slice
}

func sum_substr(my_slice []string) []string {
	var my_len = len(my_slice) - 1
	for i := 0; i < my_len; i++ {
		var x int
		var v string = my_slice[i]

		if v == "+" {
			a, _ := strconv.Atoi(my_slice[i-1])
			b, _ := strconv.Atoi(my_slice[i+1])
			x = a + b
		} else if v == "-" {
			a, _ := strconv.Atoi(my_slice[i-1])
			b, _ := strconv.Atoi(my_slice[i+1])
			x = a - b
		}

		if v == "+" || v == "-" {
			my_slice = removePart(my_slice, i-1, i)
			my_slice[i-1] = strconv.Itoa(x)
			return sum_substr(my_slice)
		}
	}

	return my_slice
}

func brackets(my_slice []string) []string {
	var my_len = len(my_slice)
	var x int
	for i := 0; i < my_len; i++ {
		var v string = my_slice[i]

		if v == "(" {
			x = i
		}

		if v == ")" {
			var temp []string = my_slice[x+1 : i]
			temp = mult_div(temp)
			temp = sum_substr(temp)
			my_slice = removePart(my_slice, x, i-1)
			my_slice[x] = temp[0]
			return brackets(my_slice)
		}
	}

	return my_slice
}

func calculate(my_str string) (res int) {
	var my_slice []string = parseStrToSlice(my_str)
	my_slice = brackets(my_slice)
	my_slice = mult_div(my_slice)
	my_slice = sum_substr(my_slice)
	res, _ = strconv.Atoi(my_slice[0])
	return
}

func main() {

	//var my_str string = "1+3*(1+2/1-(2-1))"
	var my_str string
	fmt.Scanln(&my_str)
	fmt.Println(calculate(my_str))
}

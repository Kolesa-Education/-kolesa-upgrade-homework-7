package main

import (
	"fmt"
	"os"
	"strconv"
)
// test version 

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Incorrect input!\n		example: 1+3*(1+2/1-(2-1))")
		return
	}
	fmt.Print(bracketsOperations(os.Args[1]))
	// runes := []rune(os.Args[1])
	// for i := 0; i < len(runes); i++ {

	// }

}
func bracketsOperations(str string) (int, error) {
	runes := []rune(str)
	i := 0
	
	for {

		if '0' > runes[i] || runes[i] > '9' {
			
			switch runes[i] {
			case '*', '/':
				nums, err := giveNumbers(str, i)
				if err != nil {
					return 0, err
				}
				if runes[i] == '*' {
					return nums[0] * nums[1], nil
				}
				if runes[i] == '/' {
					return nums[0] / nums[1], nil
				}
			case '+', '-':
				nums, err := giveNumbers(str, i)
				if err != nil {
					return 0, err
				}
				if runes[i] == '+' {
					return nums[0] + nums[1], nil
				}
				if runes[i] == '-' {
					return nums[0] - nums[1], nil
				}
			default:
				fmt.Println(runes[i])
				return 0, fmt.Errorf("incorrect operation\n")
			}
		}
	}

}
func giveNumbers(str string, i int) ([]int, error) {
	start := 0
	end := 0
	var arr []int
	num1, err := strconv.Atoi(str[start:i])
	if err != nil {
		return nil, fmt.Errorf("incorrect number")
	}
	arr = append(arr, num1)
	for j := i + 1; j < len(str); j++ {
		if '9' > str[j] || str[j] > '0' {
			end = j
			break
		}

	}
	num2, err := strconv.Atoi(str[i+1 : end])
	if err != nil {
		return nil, fmt.Errorf("incorrect number")
	}
	arr = append(arr, num2)
	return arr, nil
}

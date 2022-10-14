package calculators

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evalrpn(tks []string) {
	var (
		nums []float64
		x    float64
	)
	pop := func() float64 {
		n := len(nums) - 1
		t := nums[n]
		nums = nums[:n]
		return t
	}
	defer func() {
		if recover() == nil && len(nums) == 1 {
			fmt.Println(x)
		} else {
			fmt.Println("error")
		}
	}()
	for _, tk := range tks {
		switch tk {
		case "+":
			x = pop() + pop()
		case "*":
			x = pop() * pop()
		case "-":
			x = pop()
			x = pop() - x
		case "/":
			x = pop()
			x = pop() / x
		default:
			var e error
			x, e = strconv.ParseFloat(tk, 64)
			if e != nil {
				panic(0)
			}
		}
		nums = append(nums, x)
	}
}

func Calc2(input string) {
	fmt.Println("Enter expression you want to calculate")

	stdin := bufio.NewReader(os.Stdin)
	for {
		s, e := stdin.ReadString('\n')
		if e != nil {
			break
		}
		if tks := strings.Fields(s); len(tks) > 0 {
			evalrpn(tks)
		}
	}
}

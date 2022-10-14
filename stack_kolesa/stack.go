package stack_kolesa

import "errors"

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("stack is empty")
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, nil
}

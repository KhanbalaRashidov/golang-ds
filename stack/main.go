package main

import "fmt"

type Stack struct {
	stack []int
}

func NewStack() Stack {
	return Stack{stack: []int{}}
}

func (s *Stack) Push(value int) {
	s.stack = append(s.stack, value)
}

func (s *Stack) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	temp := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return temp, nil
}

func main() {

	stackValues := NewStack()
	stackValues.Push(1)
	stackValues.Push(2)
	fmt.Println(stackValues.stack)
	fmt.Println(stackValues.Pop())
	fmt.Println(stackValues.Pop())
}

package main

import "fmt"

// Stack represents a stack
type Stack []string

// IsEmpty check if stack is empty
func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

// Push append a new value onto the stack
func (stack *Stack) Push(str string) {
	*stack = append(*stack, str)
}

// Pop remove and return top element of stack
func (stack *Stack) Pop() (string, bool) {
	if stack.IsEmpty() {
		return "", false
	}

	index := len(*stack) - 1
	element := (*stack)[index]
	*stack = (*stack)[:index]
	return element, true
}

func main() {
	var stack Stack

	stack.Push("this")
	stack.Push("is")
	stack.Push("sparta!!")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}

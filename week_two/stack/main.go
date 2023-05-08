package main

import "fmt"

// Stack implementation (LIFO)

var stack = make([]string, 0)

func main() {
	push("push-example")
	if !isEmpty() {
		fmt.Println("Stack is not empty")
	}
	pop()
	if isEmpty() {
		fmt.Println("Stack is empty")
	}
}

func push(item string) {
	stack = append(stack, item)
	fmt.Printf("Pushed: %s \n", item)
}

func pop() {
	if isEmpty() {
		fmt.Println("Stack is empty")
		return
	}

	item := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Printf("Popped: %s \n", item)
}

func isEmpty() bool {
	return len(stack) == 0
}

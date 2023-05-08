package main

import "fmt"

// Queue implementation (FILO)

var queue = make([]string, 0)

func main() {
	enqueue("first")
	enqueue("second")

	if !isEmpty() {
		fmt.Println("Queue is not empty")
	}

	dequeue()
	fmt.Printf("Size of queue is: %v \n", len(queue))

	dequeue()
	if isEmpty() {
		fmt.Println("Queue is empty")
	}
}

func enqueue(item string) {
	queue = append([]string{item}, queue...)
	fmt.Printf("Enqueued: %s \n", item)
}

func dequeue() {
	if isEmpty() {
		fmt.Println("Queue is empty")
		return
	}
	item := queue[len(queue)-1]
	queue = queue[:len(queue)-1]
	fmt.Printf("Dequeued: %s \n", item)
}

func isEmpty() bool {
	return len(queue) == 0
}

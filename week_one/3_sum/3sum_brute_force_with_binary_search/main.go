package main

import (
	"fmt"
	"sort"
	"time"
)

// O(N * N * log(N))

func main() {
	numbers := []int{30, -40, -20, -10, 40, 0, 10, 5}

	start := time.Now()

	// sort makes binary search possible
	sort.Ints(numbers)
	count(numbers)
	fmt.Println("Elapsed: ", time.Since(start).Microseconds())
}

func count(numbers []int) int {
	count := 0
	for i := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			find := -(numbers[i] + numbers[j])
			if binarySearch(j+1, find, numbers) {
				fmt.Printf("%v, %v, %v = 0 \n", numbers[i], numbers[j], find)
				count++
			}
		}
	}

	return count
}

func binarySearch(low int, number int, numbers []int) bool {
	hi := len(numbers) - 1

	for {
		if low > hi {
			return false
		}

		midIndex := (low + hi) / 2
		midNumber := numbers[midIndex]

		if midNumber == number {
			return true
		}

		if midNumber > number {
			// go left
			hi = midIndex - 1
		} else {
			// go right
			low = midIndex + 1
		}
	}
}

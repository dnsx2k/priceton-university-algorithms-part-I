package main

import (
	"fmt"
	"time"
)

// Big O - cubic

func main() {
	numbers := []int{30, -40, -20, -10, 40, 0, 10, 5}

	start := time.Now()
	count(numbers)
	fmt.Println("Elapsed: ", time.Since(start).Microseconds())
}

func count(numbers []int) int {
	count := 0
	for i := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 0 {
					fmt.Printf("%v, %v, %v = 0 \n", numbers[i], numbers[j], numbers[k])
					count++
				}
			}
		}
	}

	return count
}

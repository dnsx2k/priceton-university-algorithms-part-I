package main

import "fmt"

// Quick-find defects: union too expensive, find (N array accesses)

// Quick-find [eager approach]
// Integer array id[] of size N.
// Interpretation: p and q are connected if they have the same id
// O(n)

func main() {
	ids := make([]int, 10)

	for i := 0; i < 10; i++ {
		ids[i] = i
	}

	if connected(ids, 0, 5) {
		fmt.Println("Connected before union")
	}

	union(ids, 0, 5)

	if connected(ids, 0, 5) {
		fmt.Println("Connected after union")
	}
}

// constant
// Check if p and q have the same id
func connected(ids []int, p, q int) bool {
	return ids[p] == ids[q]
}

// quadratic
// To merge components containing p and q, change all entries whose id equals id[p] to id[q]
// at most 2N + 2 array accesses
func union(ids []int, p, q int) {
	pID := ids[p]
	qID := ids[q]
	for i := range ids {
		if ids[i] == pID {
			ids[i] = qID
		}
	}
}

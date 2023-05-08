package main

import "fmt"

// Quick-union defects: trees can get tall, find too expensive (could be N array accesses)

// Quick union [lazy-approach]
// Integer array id[] of size N
// id[i] is parent of i

func main() {
	ids := make([]int, 10)
	for i := 0; i < 10; i++ {
		ids[i] = i
	}

	ids[0] = 0
	ids[1] = 1
	ids[2] = 9
	ids[3] = 4
	ids[4] = 9
	ids[5] = 6
	ids[6] = 6
	ids[7] = 7
	ids[8] = 8
	ids[9] = 9

	if connected(ids, 1, 9) {
		fmt.Println("Connected before union")
	}

	union(ids, 1, 9)

	if connected(ids, 1, 9) {
		fmt.Println("Connected after union")
	}
}

var sizes = make([]int, 10)

// Check if p and q have the same root
func connected(ids []int, p, q int) bool {
	pRoot := root(ids, p)
	qRoot := root(ids, q)

	return pRoot == qRoot
}

func root(ids []int, p int) int {
	parent := ids[p]
	if parent == ids[parent] {
		return parent
	}

	return root(ids, parent)
}

// To merge components containing p and q, set the id of p's root to the id of q's root
func union(ids []int, p, q int) {
	pRoot := root(ids, p)
	qRoot := root(ids, q)

	ids[pRoot] = qRoot
}

func union_with_weighting(ids []int, p, q int) {
	pRoot := root(ids, p)
	qRoot := root(ids, q)

	pLen := sizes[pRoot]
	qLen := sizes[qRoot]

	if pLen > qLen {
		ids[qRoot] = pRoot
		sizes[pRoot] += sizes[qRoot]
	} else {
		ids[pRoot] = qRoot
		sizes[qRoot] += sizes[pRoot]
	}
}

func root_with_path_compression(ids []int, p int) int {
	for {
		if p == ids[p] {
			break
		}

		ids[p] = ids[ids[p]]
		p = ids[p]
	}
	return p
}

func union_with_path_compression(ids []int, p, q int) {
	pRoot := root_with_path_compression(ids, p)
	qRoot := root_with_path_compression(ids, q)

	pLen := sizes[pRoot]
	qLen := sizes[qRoot]

	if pLen > qLen {
		ids[qRoot] = pRoot
		sizes[pRoot] += sizes[qRoot]
	} else {
		ids[pRoot] = qRoot
		sizes[qRoot] += sizes[pRoot]
	}
}

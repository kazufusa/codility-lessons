package main

import (
	"fmt"
	"testing"
)

func NumberSolitaire(A []int) int {
	b := make([]int, len(A))
	b[0] = A[0]
	for i := 1; i < len(A); i++ {
		score := b[i-1]
		for j := 1; j < 6; j++ {
			if j < i && score < b[i-j-1] {
				score = b[i-j-1]
			}
		}
		b[i] = score + A[i]
	}
	return b[len(b)-1]
}

func TestNumberSoliteire(t *testing.T) {
	A := []int{1, -2, 0, 9, -1, -2}
	fmt.Println(NumberSolitaire(A))
	A = []int{-2, 5, 1}
	fmt.Println(NumberSolitaire(A))
	A = []int{-2, -1, 1}
	fmt.Println(NumberSolitaire(A))
}

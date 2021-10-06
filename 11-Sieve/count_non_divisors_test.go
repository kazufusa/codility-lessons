package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CountNonDivisors(A []int) []int {
	N := len(A)
	ret := make([]int, N)
	counts := make(map[int]int)
	divisors := make(map[int][]int)
	var ok bool
	var ds []int

	for _, a := range A {
		if _, ok = counts[a]; ok {
			counts[a]++
		} else {
			counts[a] = 1
		}
	}

	for i, a := range A {
		ret[i] = N
		if ds, ok = divisors[a]; !ok {
			ds = Divisors(a)
		}
		for _, d := range ds {
			ret[i] -= counts[d]
		}
	}

	return ret
}

func Divisors(x int) []int {
	ret := []int{}
	_x := int(math.Sqrt(float64(x)))
	for i := 1; i < _x; i++ {
		if x%i == 0 {
			ret = append(ret, i, x/i)
		}
	}
	if x%_x == 0 {
		ret = append(ret, _x)
		if _x*_x != x {
			ret = append(ret, x/_x)
		}
	}
	return ret
}

func TestCountNonDivisors(t *testing.T) {
	assert.Equal(t, []int{2, 4, 3, 2, 0}, CountNonDivisors([]int{3, 1, 2, 3, 6}))
}

package main

import (
	"fmt"
	"testing"
)

// Sieve returns prime numbers
func Sieve(n int) []bool {
	ret := make([]bool, n+1)
	for i := 2; i < n; i++ {
		ret[i] = true
	}
	i := 2
	for i*i <= n {
		if ret[i] {
			k := i * i
			for k <= n {
				ret[k] = false
				k += i
			}
		}
		i++
	}

	return ret
}

func TestSieve(t *testing.T) {
	for i, v := range Sieve(20) {
		fmt.Println(i, v)
	}
}

func ArrayF(n int) []int {
	ret := make([]int, n+1)
	i := 2
	for i*i <= n {
		if ret[i] == 0 {
			k := i * i
			for k <= n {
				if ret[k] == 0 {
					ret[k] = i
				}
				k += i
			}
		}
		i++
	}

	return ret
}

func TestArrayF(t *testing.T) {
	for i, v := range ArrayF(20) {
		fmt.Println(i, v)
	}
}

func Factorize(x int, F []int) []int {
	var ret []int
	for F[x] > 0 {
		ret = append(ret, F[x])
		x /= F[x]
	}
	ret = append(ret, x)
	return ret
}

func TestFactorize(t *testing.T) {
	fmt.Println(Factorize(20, ArrayF(20)))
}

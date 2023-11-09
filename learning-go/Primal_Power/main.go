package main

import (
	"fmt"
	"math"
)

func main() {

	a := []int{-6, 10, 12,11}
	fmt.Println(solve(a))
}
func solve(A []int) int {
	c := 0

	for _, v := range A {
		if v <= 1 {
			continue
		} else {
			var y bool
			for i := 2; i <= int(math.Sqrt(float64(v))); i++ {
				fmt.Println(v,int(math.Sqrt(float64(v))))
				if v%i == 0 {
					y = true
				}
			}
			if !y {
				c++
			}
		}

	}

	return c
}

func findPrime(A []int) int {
	c := 0
	//res := make(map[int]int)
	if len(A) < 1 {
		return 0
	} else if len(A) == 1 && A[0] == 1 {
		return 1
	} else {
		for _, v := range A {
			if v <= 1 {
				continue
			} else {
				var y bool
				for i := 2; i <= int(math.Sqrt(float64(v))); i++ {
					fmt.Println(v, i, v%i)
					if v%i == 0 {
						y = true
					}
				}
				if !y {
					c++
				}
			}
		}
	}
	return c
}

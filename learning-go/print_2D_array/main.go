package main

import "fmt"

func main() {
	fmt.Println(solve(4))
}
func solve(n int) [][]int {
	res := make([][]int, n)
	for i := 1; i <=n; i++ {
		t := make([]int, n)
		for j := 0; j < i; j++ {
			t[j] = j + 1
		}
		res[i-1] = t
	}
	return res
}

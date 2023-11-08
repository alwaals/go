package main

import "fmt"

// Given an array A and an integer B.
// A pair(i, j) in the array is a good pair if i != j and (A[i] + A[j] == B). Check if any good pair exist or not.
func main() {
	i:=[]int{1,2,3}
	fmt.Println(findPair(i,7))
}
func findPair(inp []int, n int) int {
	for i := 0; i < len(inp); i++ {
		for j := 0; j < len(inp); j++ {
			if inp[i]+inp[j] == n && i!=j {
				return 1
			}
		}
	}
	return 0
}

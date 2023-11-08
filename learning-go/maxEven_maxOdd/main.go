package main

import (
	"fmt"
	"math"
)

func main(){

	i := []int{5, 17, 100, 1}
	fmt.Println(solve(i))

}

func solve(inp []int) int{

	//diff := 0
	minOdd,maxEven := math.MaxInt64,math.MinInt64
	for _,v:=range inp{
		if v%2!=0{
			minOdd = findMin(minOdd,v)
		}else{
			maxEven = findMax(maxEven,v)
		}
	}
	return maxEven-minOdd
}

func findMin(i,j int) int{
	if i>j{return j}
	return i
}

func findMax(i,j int) int{
	if i>j{return i}
	return j
}
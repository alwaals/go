package main

import "fmt"

func main(){
	A:=[]int{343,291,963,165,152}
	fmt.Println(findSubarray(A))
}

func findSubarray(A []int) []int{
	min:=A[0]
	max:=A[0]
	minIndex,maxIndex:=0,0
	for i:=1;i<len(A);i++{
		if A[i]>max{
			max=A[i]
			maxIndex=i
		}
		if A[i]<min{
			min=A[i]
			minIndex=i
		}
	}
	fmt.Println(minIndex,min,maxIndex,max)
	if maxIndex<minIndex{
		return A[maxIndex:minIndex+1]
	}
	return A[minIndex:maxIndex+1]
}

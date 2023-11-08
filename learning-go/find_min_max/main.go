package main

import (
	"fmt"
)

func main(){

     arr := []int{2,3,4,7}
	 fmt.Print(solve(arr))
}

func solve(inp []int) int{

	res:=0
	minV, maxV := inp[0],inp[len(inp)-1]
	for _,i :=range inp{
		minV = min(minV,i)
		maxV = max(maxV,i)
	}
	fmt.Println(minV,maxV)
	for _,i := range inp{
		fmt.Println("For i:",i)
		if i!=minV && i!=maxV{
			res=res+1
		}
	}
	return res
}
func min(i,v int)int{
 if i<v{return i}
 return v
}

func max(i,v int)int{
	if i>v{return i}
	return v
}
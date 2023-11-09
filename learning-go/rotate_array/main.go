package main

import "fmt"

func main(){
i:=[]int{1, 2, 3, 4, 5,1}
fmt.Println(rotate(i,3))
}

func rotate(inp []int,b int)[]int{
	var res []int
	// lastEle := inp[b:]
	// res = inp[0:len(lastEle)]
	// fmt.Println(lastEle,res)

	secondHalf := inp[0:len(inp)-b]
	res=inp[len(inp)-b:]
	//fmt.Println(secondHalf,res)
	res= append(res, secondHalf...)
	return res
}
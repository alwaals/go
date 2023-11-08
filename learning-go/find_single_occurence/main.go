package main

import "fmt"

func main(){

	i := []int{1,3,1,3,2}
	fmt.Println(solve(i))
}

func solve(inp []int) int{

	res:=inp[0]

	for i:=1;i<len(inp);i++{
		//fmt.Println(res^inp[i])
		fmt.Println(res,inp[i])
		res=res^inp[i]//1,1,0,
		fmt.Println(res)
	}
	return res
}
package main

import "fmt"

func main(){

	inp := []int{1,2,4,6,8}
	fmt.Println(solve(inp))

}

func solve(inp []int) []int{

	res := []int{}

	for i:=0;i<len(inp);i++{
		c:=0
		for j:=0;j<len(inp);j++{
			if inp[i]<inp[j] && i!=j{
				c+=1
			}
			if c>=2{
				res=append(res, inp[i])
				break
			}
		}
	}
	return res
}
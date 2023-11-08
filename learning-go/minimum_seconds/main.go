package main

import "fmt"

func main(){

	i := []int{2, 4, 1, 3, 2}
	max,sum:=maxAndSum(i)
	fmt.Println(max,sum,4*len((i))-sum)
}

func maxAndSum(inp []int) (int,int){
	max:=inp[0]
	sum:=inp[0]
	for _,v:=range inp[1:]{
		sum+=v
		if v>max{
			max=v
		}
	}
	return max,sum
}
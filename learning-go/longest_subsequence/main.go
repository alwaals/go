package main

import "fmt"

//Given an array of integers A of size, N. Find the longest subsequence of A, which is odd-even.
func main(){
	i:=[]int{1, 2, 2, 5, 6,7,8,6}
	fmt.Println(findLongest(i))
}
func findLongest(inp []int) int{
	var res []int 
	evenBool,oddBool:=false,false
	if inp[0]%2==0{
		evenBool=false
		oddBool=true
	}else{
		evenBool=true
		oddBool=false
	}
	res=append(res, inp[0])
	for i:=1;i<len(inp);i++{
	//    fmt.Println("Num is ",inp[i],evenBool,oddBool)
	   switch{
	   case inp[i]%2==1&&oddBool:
		evenBool=true
		oddBool=false
		res=append(res, inp[i])
	   case inp[i]%2==0&&evenBool:
		evenBool=false
		oddBool=true
		res=append(res, inp[i])
	   }
	}
	fmt.Println(res)
	return len(res)
}
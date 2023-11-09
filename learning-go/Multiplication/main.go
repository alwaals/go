package main

import "fmt"

func main(){

	//A:=[]int{1, 2, 3, 4, 5}
	A:=[]int{5, 17, 100, 11}
	fmt.Println(multiply(A))
}

func multiply(A []int)[]int{
    var res []int
	if len(A)==1{
        return A
    }
	for i:=0;i<len(A);i++{
		if i==0{
			res =append(res, A[i]*A[i+1])
		}else if i==len(A)-1{
			res =append(res, A[i]*A[i-1])
		}else{
			res =append(res, A[i+1]*A[i-1])
		}
	}
	return res
}
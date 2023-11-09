package main

import "fmt"

func main(){
	A := "ABCGAG"
	fmt.Println(findSubstr(A))

}
func findSubstr(str string) int{
	ag,g:=0,0
	for i:=len(str)-1;i>=0;i--{
		if str[i]=='G'{
			g++
		}
		if str[i]=='A'{
			ag=ag+g
		}
	}
	return ag
}
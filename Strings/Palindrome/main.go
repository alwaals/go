package main

import "fmt"

func main(){
	fmt.Println(palindrome("nitin"))
}
func palindrome(s1 string) string{
	lastInd := len(s1)-1
	for i:=0;(i<len(s1)/2)&&(i<lastInd-1);i++{
		if s1[i] != s1[lastInd-i]{
			return fmt.Sprintf("Not a Palindrome")
		}
	}
   return fmt.Sprintf("Its a palindrome")
}
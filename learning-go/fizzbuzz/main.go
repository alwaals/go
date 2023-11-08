package main

import "fmt"

//Given a positive integer A, return an array of strings with all the integers from 1 to N. But for multiples of 3
//the array should have “Fizz” instead of the number. For the multiples of 5, the array should have “Buzz”
//instead of the number. For numbers which are multiple of 3 and 5 both, the array should have "FizzBuzz" instead of the number.
func main(){
	find(50)
}

func find(n int){
	for value:=1;value<=n;value++{
	switch {
	case value%5==0 && value%3==0:
		fmt.Println(value,"FizzBuzz")
	case value%5==0:
		fmt.Println(value,"Buzz")
	case value%3==0:
		fmt.Println(value,"Fizz")
	default:
     	fmt.Println(value)
	}
}
}
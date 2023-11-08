package main

import "fmt"

func main(){

	sl := []int{1,0,5,0,7}
	fmt.Println("Before:",sl)

	// //Move 0's to the End
	// for i:=0;i<len(sl);i++{
	// 	if sl[i]==0{
	// 		sl = append(sl[0:len(sl)])
	// 	}
	// }
	// Remove 1st element
	sl = append(sl[1:len(sl)])
	fmt.Println("Before:",sl)

	//Remove last element
	sl = append(sl[0:len(sl)-1])
	fmt.Println(sl)

}
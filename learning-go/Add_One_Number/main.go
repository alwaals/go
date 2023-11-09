package main

import "fmt"

func main() {
	i := []int{1, 9, 9,9,9}
	fmt.Println(addNumber(i))
}
func addNumber(inp []int) []int {
	borrow:=false
	if len(inp)>1{
	for i := len(inp) - 1; i > 0; i-- {
		if inp[i]+1 < 10 && !borrow{
			inp[i] = inp[i] + 1
			break
		} else if !borrow{
			inp[i] = 0
			if inp[i]+1 < 10 {
				inp[i-1] = inp[i-1] + 1
				if inp[i-1]>9{
					borrow=true
					continue
				}
				break
			} else {
				borrow=true
			}
		} else {
			if inp[i] < 10{
				break
			}else {
				inp[i] = 0
				inp[i-1] = inp[i-1] + 1
				borrow=true
			}
		}
	}
}else{
	inp[0]=inp[0]+1
}
	return inp
}
func addition(inp []int) {
}

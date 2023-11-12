package main

import "fmt"

func main(){
fmt.Println(anagram("silent","litten"))
}
func anagram(s1 string,s2 string) string{
	lenS1:=len(s1)
	lenS2:=len(s2)

	if lenS1 != lenS2{
		return fmt.Sprintf("Not anagrams")
	}

	mp := make(map[string]int)

	for i:=0;i<lenS1;i++{
		mp[string(s1[i])]++
	}
	fmt.Println("Before:",mp)

	for i:=0;i<lenS2;i++{
		fmt.Println(s1[i])
		mp[string(s1[i])]-- 
		fmt.Println("After:",mp)
	}

	fmt.Println("After:",mp)
	for _,v := range mp{
		if v>0{
			return fmt.Sprintf("Not anagrams")
		}
	}
	return fmt.Sprintf("Anagrams")
}
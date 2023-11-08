package main

import "fmt"
func main(){

	mp := make(map[string]string)
	mp["CA"]="California"
	mp["TX"]="Texas"
	mp["SE"]="Seattle"
	mp["LV"]="Los Vegas"

	fmt.Println(len(mp))
	for k,v:=range mp{
		fmt.Println(fmt.Sprintf("Value at index %s:%s",k,v))
	}
	//Delete particular value from Map
	//delete(mp,"TX")

	//Remove first element
	
}
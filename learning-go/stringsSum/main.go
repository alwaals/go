// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {
   
	f1,err:=strconv.ParseFloat(strings.TrimSpace(value1),64)
	if err!=nil{
		return 0
	}

	f2,err:=strconv.ParseFloat(strings.TrimSpace(value2),64)
	if err!=nil{
		return 0
	}
	return f1+f2
}
func main(){
	fmt.Print(calculate("10.3","30.2"))
}

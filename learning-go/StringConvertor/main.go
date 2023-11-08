package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text:")
	inp, _ := reader.ReadString('\n')
	fmt.Print("Entered text is :", inp)
	fmt.Print("Enter number:")
	inpN, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(inpN), 64)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print("Number is:", aFloat)
	}
}

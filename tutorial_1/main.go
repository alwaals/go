package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main(){
fmt.Println("Hello Tutorial_1")
var ch = make(chan int,5)
var exit = make(chan struct{})
for i:=0;i<5;i++{
	wg.Add(1)
    go readChannel(ch)
	//fmt.Println("Reading chan",<-ch)
}
wg.Wait()
close(ch)
for {
	select {
	case i,open:=<-ch:
		if !open{
			fmt.Println("Exiting readChannel:")	
			close(exit)
			break
		}
		fmt.Println("Reading from ch:",i)
	}
}
}
func readChannel(ch chan int){
	defer wg.Done()
	time.Sleep(2000)
	for i:=0;i<5;i++{
		ch <- i	
	}
}
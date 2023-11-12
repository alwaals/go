package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	t := time.Now()

	ch := make(chan string,2)
	exit := make(chan struct{})
	var wg sync.WaitGroup
	user:=getUser()
	wg.Add(2)
	go appendString(user,ch,&wg)
	go appendInt(user,ch,&wg)
		wg.Wait()
		close(ch)	
	// for i:=range ch{
	// 	fmt.Println("Got from channel:",i)
	// }
	for{
		select {
		case i,open:=<-ch:
			if !open{
				fmt.Println("Total time taken:",time.Since(t))
				close(exit)
				return
			}
			fmt.Println("Got from channel:",i)
		}
	}
}
func getUser() string{
	time.Sleep(time.Millisecond*100)
	return "Bob"
}
func appendString(user string,ch chan string,wg *sync.WaitGroup){
	time.Sleep(time.Millisecond*100)
	ch <- fmt.Sprintf("Hi:%s",user)
	wg.Done()
	return
}
func appendInt(user string,ch chan string,wg *sync.WaitGroup){
	time.Sleep(time.Millisecond*100)
	ch <- fmt.Sprintf("11:%s",user)
	wg.Done()
	return
}
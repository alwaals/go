package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type transactionProcessor struct {
	Datetime  string `json:"datetime"`
	Partition string `json:"partition"`
	Value     string `json:"value"`
}

var mut sync.RWMutex
var wg sync.WaitGroup

func main() {
	trans := []transactionProcessor{}
	ch := make(chan []byte)

	wg.Add(1)
	go callGoroutine(ch)
	// defer close(ch)

	for {
		select {
		case out := <-ch:
			fmt.Println(string(out))
			var t transactionProcessor
			err := json.NewDecoder(bytes.NewBuffer(out)).Decode(&t)
			ifError(err)
			trans = append(trans, t)
		default:
			break
		}
	}
	wg.Wait()
	fmt.Println(len(trans))
}

func callGoroutine(ch chan<- []byte) {
	f := OpenFile()
	defer f.Close()
	defer wg.Done()
	defer close(ch)
	// for i:=0;i<5;i++{
	// 	ch<-i
	// }
	b := bufio.NewReader(f)
	//c := []byte{}
	for {
		str, _, err := b.ReadLine()
		if err != nil {
			if io.EOF == err {
				break
			}
			log.Fatal(err.Error())
		}
		//fmt.Println(string(str))
		ch <- str
	}
}

func OpenFile() *os.File {
	file, err := os.Open("data.json")
	ifError(err)
	return file
}
func ifError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

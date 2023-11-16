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
	"time"
)

type transaction_processor struct {
	Datetime  string `json:"datetime"`
	Value     string `json:"value"`
	Partition string `json:"partition"`
}

func OpenFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Got error during opening file:%s", err.Error())
	}
	return file, nil
}
func ReadingFileUnmarshalling(file *os.File) ([]transaction_processor, error) {
	var trans []transaction_processor
	bufReader := bufio.NewReader(file)
	for {
		str, _, err := bufReader.ReadLine()
		if err != nil {
			if io.EOF == err {
				break
			}
			return trans, fmt.Errorf("Got error during ReadString:%s", err.Error())
		}
		//reader := bytes.NewBuffer(str)
		var transaction transaction_processor
		err = json.Unmarshal(str, &transaction)
		if err != nil {
			return nil, err
		}
		trans = append(trans, transaction)
	}
	return trans, nil
}
func ReadingFileDecoder(file *os.File) ([]transaction_processor, error) {
	var trans []transaction_processor
	bufReader := bufio.NewReader(file)
	for {
		str, _, err := bufReader.ReadLine()
		if err != nil {
			if io.EOF == err {
				break
			}
			return trans, fmt.Errorf("Got error during ReadString:%s", err.Error())
		}
		reader := bytes.NewBuffer(str)
		var transaction transaction_processor
		err = json.NewDecoder(reader).Decode(&transaction)
		if err != nil {
			return nil, err
		}
		trans = append(trans, transaction)
	}
	return trans, nil
}
func startWorker(lines <-chan string) <-chan string {
    finished := make(chan string)
    go func() {
        defer close(finished)
        for line := range lines {
            // Do your heavy work here
			fmt.Println("line:",line)
            finished <- line
        }
    }()
    return finished
}

func merge(cs ...<-chan string) <-chan string {
    var wg sync.WaitGroup
    out := make(chan string)
    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed, then calls wg.Done.
    output := func(c <-chan string) {
        for n := range c {
        out <- n
        }
        wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // Start a goroutine to close out once all the output goroutines are
    // done.  This must start after the wg.Add call.
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
func main() {
	file, err := OpenFile("data.json")
	if err != nil {
		fmt.Println(fmt.Errorf("Got error during opening file:%s", err.Error()))
		return
	}
	defer file.Close()

	t := time.Now()
	var trans []transaction_processor

	// transactions,err:=ReadingFileUnmarshalling(file)
	// if err!=nil{
	// 	fmt.Println(fmt.Errorf("Got error during ReadingFile:%s",err.Error()))
	// 	return
	//  }
	//fmt.Println("Total length of transactions:",len(transactions))

	lines := make(chan string)
	// start four workers to do the heavy lifting
	wc1 := startWorker(lines)
	wc2 := startWorker(lines)
	wc3 := startWorker(lines)
	wc4 := startWorker(lines)
	scanner := bufio.NewScanner(file)
	go func() {
		defer close(lines)
		for scanner.Scan() {
			lines <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	merged:=merge(wc1, wc2, wc3, wc4)
	for line := range merged {

		//fmt.Println(fmt.Sprintf("%T", line))
		var transaction transaction_processor
		reader := bytes.NewBuffer([]byte(line))
		err = json.NewDecoder(reader).Decode(&transaction)
		if err != nil {
			log.Fatal(err)
		}
		trans = append(trans, transaction)
	}
	fmt.Println("Length of the file:",len(trans))
	fmt.Println("Total time taken:", time.Since(t))
}

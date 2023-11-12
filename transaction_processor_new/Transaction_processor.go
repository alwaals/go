package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type transaction_processor struct {
	Datetime  string `json:"datetime"`
	Value     string `json:"value"`
	Partition string `json:"partition"`
}

func OpenFile(fileName string) (*os.File,error){
	file,err:=os.Open(fileName)
	if err!=nil{
	   return nil,fmt.Errorf("Got error during opening file:%s",err.Error())
	}
	return file,nil
}
func ReadingFileUnmarshalling(file *os.File) ([]transaction_processor,error){
	var trans []transaction_processor
	bufReader:=bufio.NewReader(file)
	for{
		str,_,err:=bufReader.ReadLine()
		if err!=nil{
			if io.EOF==err{
				break
			}
			return trans,fmt.Errorf("Got error during ReadString:%s",err.Error())
		 }
		//reader := bytes.NewBuffer(str)
		var transaction transaction_processor
		err = json.Unmarshal(str,&transaction)
		if err != nil {
				return nil, err
		}
		trans = append(trans, transaction)
		}		
	return trans,nil
}
func ReadingFileDecoder(file *os.File) ([]transaction_processor,error){
	var trans []transaction_processor
	bufReader:=bufio.NewReader(file)
	for{
		str,_,err:=bufReader.ReadLine()
		if err!=nil{
			if io.EOF==err{
				break
			}
			return trans,fmt.Errorf("Got error during ReadString:%s",err.Error())
		 }
		reader := bytes.NewBuffer(str)
		var transaction transaction_processor
		err = json.NewDecoder(reader).Decode(&transaction)
		if err != nil {
				return nil, err
		}
		trans = append(trans, transaction)
		}		
	return trans,nil
}
func main() {
     f,err:=OpenFile("data.json")
	 if err!=nil{
		fmt.Println(fmt.Errorf("Got error during opening file:%s",err.Error()))
		return
	 }
	defer f.Close()

	t:=time.Now()

	transactions,err:=ReadingFileUnmarshalling(f)
	if err!=nil{
		fmt.Println(fmt.Errorf("Got error during ReadingFile:%s",err.Error()))
		return
	 }
	fmt.Println("Total length of transactions:",len(transactions))
	fmt.Println("Total time taken:",time.Since(t))
}

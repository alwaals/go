package main

import (
	"bufio"
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
func ReadingFile(file *os.File) ([]transaction_processor,error){
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
		 //fmt.Println(string(str))
		 var transaction transaction_processor
		 errM:=json.Unmarshal(str,&transaction)
		 if err!=nil{
			return trans,fmt.Errorf("Got error during unMarshalling:%s",errM.Error())
		  }
		  trans = append(trans, transaction)
	}
	return trans,nil
}
func main() {
     file,err:=OpenFile("data.json")
	 if err!=nil{
		fmt.Println(fmt.Errorf("Got error during opening file:%s",err.Error()))
		return
	 }
	defer file.Close()
	t:=time.Now()
	trans,errT:=ReadingFile(file)
	if errT!=nil{
		fmt.Println(fmt.Errorf("Got error reading file:%s",errT.Error()))
		return
	 }
	for _,v:=range trans{
		fmt.Println(fmt.Sprintf("%+v",v))
	}
	fmt.Println("Length of the file:",len(trans))
	fmt.Println("Total time taken:",time.Since(t))
}

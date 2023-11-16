package main

import (
	"fmt"
	"net/http"
)

func main(){
 fmt.Println("Started service!")
 mux:=http.NewServeMux()
 mux.HandleFunc("/hello",hello)
 http.ListenAndServe(":8080",mux)

//  client:= http.Client{
// 	Timeout: 20 * time.Second,
//  }
//  client.Do()
}
func hello(w http.ResponseWriter,req *http.Request){

	w.Write([]byte("Here is a hello string...."))
	return
}

type Shape struct{
	Name string
	Height float64
	Width float64
}

type GetInfo interface{
    area()
	perim()
}

func (s *Shape) area(){
	fmt.Println(fmt.Sprintf("Area of %s:",s.Name))
}


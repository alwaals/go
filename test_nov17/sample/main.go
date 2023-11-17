package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main(){
	fmt.Println("Starting service")
	

	mux:=http.NewServeMux()
	mux.HandleFunc("/hello",handleFunc)
	newLogger := NewLogger(mux)
	http.ListenAndServe(":8080",newLogger)
}
func handleFunc(w http.ResponseWriter,req *http.Request){
	// Open a JPG file.
    f, _ := os.Open("/Users/swetha/Downloads/flower.webp")

    // Read the entire JPG file into memory.
    reader := bufio.NewReader(f)
    content, _ := ioutil.ReadAll(reader)

    // Set the Content Type header.
    w.Header().Set("Content-Type", "image/jpeg")

    // Write the image to the response.
    w.Write(content)
	return
}
type Logger struct{
	http http.Handler
}
func NewLogger(h http.Handler) *Logger{
	return &Logger{h}
}
func (l *Logger) ServeHTTP(w http.ResponseWriter,req *http.Request){
	t := time.Now()
    l.http.ServeHTTP(w,req)
	if req.Method == "POST"{
       fmt.Printf("Accepts only GET methods\n")
	   return
	}
	fmt.Printf("%s,%s,%s\n",req.Method,req.URL.Path,time.Since(t))
}
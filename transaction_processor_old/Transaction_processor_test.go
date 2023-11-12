package main

import (
	"fmt"
	"strings"
	"testing"
)

// Expecting failure since fileName not provided
func Test_OpenFile_SuccessCase(t *testing.T) {
	_,got:= OpenFile("data.json")
    if got != nil {
        t.Errorf("got %q, wanted nil", got)
    }
}

func Test_OpenFile_FailureCase(t *testing.T) {
	_,got:= OpenFile("")
	want := fmt.Errorf("Got error during opening file")

    if !strings.Contains(string(got.Error()),"Got error") {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func Test_ReadingFile_SuccessCase(t *testing.T){
	file,_:= OpenFile("data.json")
	_,err:=ReadingFile(file)
	fmt.Println(err)
}

func Test_ReadingFile_FailureCase1(t *testing.T){
	_,got:=ReadingFile(nil)
	want := fmt.Errorf("Got error during ReadString:invalid argument")

    if !strings.Contains(string(got.Error()),"Got error during ReadString") {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func Test_ReadingFile_FailureCase2(t *testing.T){
	file,_:= OpenFile("")
	_,got:=ReadingFile(file)
	fmt.Print(got)
	want := fmt.Errorf("Got error during ReadString:invalid argument")

    if !strings.Contains(string(got.Error()),"Got error during ReadString") {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

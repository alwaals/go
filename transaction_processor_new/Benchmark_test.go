package main

import "testing"


func BenchmarkReadingFileUnmarshalling(b *testing.B) {
	file,_:=OpenFile("data.json")
	for n := 0; n < b.N; n++ {
       ReadingFileUnmarshalling(file)
	}
}

func BenchmarkReadingFileDecoder(b *testing.B) {
	file,_:=OpenFile("data.json")
	for n := 0; n < b.N; n++ {
       ReadingFileDecoder(file)
	}
}
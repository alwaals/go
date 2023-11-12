package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results []string
var mut sync.Mutex
var m sync.RWMutex

func main() {

	t := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go count(i)
	}
	wg.Wait()
	fmt.Println("Total time taken:", time.Since(t))
	fmt.Println("The result of database is : ", results)
}
func count(i int) {
	defer wg.Done()
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()

}
func save(str string) {
	mut.Lock()
	results = append(results, str)
	mut.Unlock()
}
func log() {
	m.RLock()
	fmt.Println("Reading results:", results)
	m.RUnlock()
}

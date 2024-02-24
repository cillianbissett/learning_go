package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// NB Concurrency != Parallelism, difference is flicking back and forth vs same time
// can get parallel execution in some cases, but not always
// there is also read/write locks, channels, and other ways to manage concurrency
var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8", "id9", "id10"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		go wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("Time taken: ", time.Since(t0))
}

func dbCall(i int) {
	// simulate a delay/lengthy query
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Result from DB: ", dbData[i])
	// need to apply a lock to stop concurrent shenanigans
	m.Lock()
	results = append(results, dbData[i])
	//get rid of lock to allow next guy in line to append
	wg.Done()
}

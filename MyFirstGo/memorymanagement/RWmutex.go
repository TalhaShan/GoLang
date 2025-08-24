/*
sync.RWMutex

Sometimes we have more readers than writers.
Multiple goroutines can read safely at the same time (no data corruption).
But writes must be exclusive (no one else can read/write while writing).

Operations
RLock() → acquire a read lock (many goroutines can hold this simultaneously).
RUnlock() → release a read lock.
Lock() → acquire an exclusive write lock (blocks both readers & writers).
Unlock() → release the write lock.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var data = 0
var rw sync.RWMutex

func read(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rw.RLock() // shared lock
	fmt.Printf("Reader %d: reading value %d\n", id, data)
	time.Sleep(100 * time.Millisecond)
	rw.RUnlock()
}

func write(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rw.Lock() // exclusive lock
	data++
	fmt.Printf("Writer %d: writing value %d\n", id, data)
	time.Sleep(200 * time.Millisecond)
	rw.Unlock()
}

func main() {
	var wg sync.WaitGroup

	// 2 writers + 3 readers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go read(i, &wg)
	}

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go write(i, &wg)
	}

	wg.Wait()
}

/*
🔹 1. Why Mutex?
When multiple goroutines share the same memory/data, they may interleave reads and writes in unsafe ways → race conditions.
Example: two goroutines incrementing a counter at the same time could overwrite each other → wrong result.
👉 To fix this, we need synchronization primitives like Mutex.
🔹 2. sync.Mutex
A Mutex (mutual exclusion lock) ensures only one goroutine at a time can access a critical section (piece of code working on shared data).
Operations
Lock() → acquire the lock (block if someone else has it).
Unlock() → release the lock so others can use it.
Example (without mutex → race condition)
*/
package main

import (
	"fmt"
	"sync"
)

var counter = 0

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // NOT safe, race condition
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}

//👉 You might expect 1000, but you’ll often see lower numbers because goroutines overwrite each other.

//Example (with sync.Mutex)
/*
var counter = 0
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()         // acquire lock
			counter++         // safe now
			mu.Unlock()       // release lock
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
*/

// ✅ Now you’ll always get 1000.

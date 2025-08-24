package main

Perfect 👌 — deadlocks and channels are key to really *understanding* Go’s concurrency model. Let’s go step by step with **theory + practical code examples**.

---

// 🔹 1. What is a Deadlock?

* A **deadlock** happens when goroutines are waiting for each other forever, so **no one can proceed**.
* In Go, you’ll often see:

```
  fatal error: all goroutines are asleep - deadlock!
  ```

if your program reaches a state where goroutines are stuck.

---

// 🔹 2. Deadlock with Mutex

Classic case: two goroutines lock resources in **opposite order**.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var mu1, mu2 sync.Mutex

func main() {
	go func() {
		mu1.Lock()
		defer mu1.Unlock()
		time.Sleep(100 * time.Millisecond)

		mu2.Lock()
		defer mu2.Unlock()
		fmt.Println("Goroutine 1 finished")
	}()

	go func() {
		mu2.Lock()
		defer mu2.Unlock()
		time.Sleep(100 * time.Millisecond)

		mu1.Lock()
		defer mu1.Unlock()
		fmt.Println("Goroutine 2 finished")
	}()

	time.Sleep(2 * time.Second) // let goroutines run
	fmt.Println("Main finished")
}
```

👉 Both goroutines are stuck:

* Goroutine 1 holds `mu1`, waits for `mu2`.
* Goroutine 2 holds `mu2`, waits for `mu1`.
* **Deadlock!**

---

// 🔹 3. Deadlock with Channels

Channels can also deadlock if you send/receive incorrectly.

////// Example: Deadlock on send (no receiver)

```go
package main

func main() {
	ch := make(chan int)
	ch <- 42 // BLOCKS forever, no one receives
}
```

////// Example: Deadlock on receive (no sender)

```go
package main

func main() {
	ch := make(chan int)
	<-ch // BLOCKS forever, no one sends
}
```

---

// 🔹 4. Solving Deadlocks with Channels

Go channels are powerful for synchronization **if used correctly**.

////// Example: Goroutines with channel communication

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
}
```

✅ No deadlock — goroutine sends, main receives.

---

////// Example: Buffered Channels

Buffered channels can store values temporarily, avoiding deadlocks when no immediate receiver exists.

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffer size 2

	ch <- 1
	ch <- 2
	fmt.Println("Sent 2 values without blocking")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

✅ Buffer allows sends without immediate receive.

---

// 🔹 5. Select Statement (avoiding deadlocks)

Use `select` to wait on multiple channels and avoid blocking forever.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout!")
		}
	}
}
```

✅ Select prevents deadlock by handling whichever channel responds first (or timing out).

---

// 🔑 Key Takeaways

* **Deadlock happens** when goroutines wait for each other forever (via Mutex or channel).
* With **Mutex**: avoid circular locking.
* With **Channels**: ensure every send has a receiver, every receive has a sender.
* Use **buffered channels** or **`select` with timeout** to avoid indefinite blocking.

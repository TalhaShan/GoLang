package main

import "fmt"

func main() {
	go greeting("Hello")
	greeting("World")
}

func greeting(s string) {

	for i := 0; i < 10; i++ {
		fmt.Println(s)
	}
}

//Ok the behavior was sometimes it prints sometime not
/* Reason ***
Perfect — let’s dig into **why** this happens in Go.

---

### 1. What happens when you call a function normally

```go
greeting("World")
```

* Runs in the **main goroutine** (Go always starts your program in a single goroutine called `main`).
* The function must **finish** before the next line runs.

---

### 2. What happens when you call with `go`

```go
go greeting("Hello")
```

* This launches a **new goroutine** — think of it like a very lightweight thread managed by the Go runtime.
* The line starts running `greeting("Hello")`, **but asynchronously**.
* `main` doesn’t wait for it; it immediately continues to the next line.

---

### 3. The special thing about `main`

* Your program **ends when the main goroutine (the one running `main()`) ends**.
* It doesn’t matter if other goroutines are still running — the process exits and kills them.
* That’s why `"Hello"` sometimes doesn’t show: the `main` function finishes after printing `"World"`, before `"Hello"` had time to run.

---

### 4. Why it sometimes prints and sometimes not

* Goroutines are scheduled by the Go runtime.
* Sometimes the `"Hello"` goroutine gets CPU time before `main` exits → you see some `"Hello"` printed.
* Sometimes `main` finishes too quickly → program ends, `"Hello"` never runs.

This is **race scheduling** — non-deterministic by design.

---

### 5. How Go solves this

To make sure goroutines finish before `main` exits, you need a **synchronization mechanism**:

* **`sync.WaitGroup`** → wait until goroutines finish.
* **channels** → send/receive signals when work is done.
* (for learning only) **`time.Sleep`** → delay `main` long enough, but that’s brittle.
---
🔑 **Core concept**:
`go` launches work **concurrently**, but **Go does not wait** for that work unless you explicitly tell it to.
`main` exiting is the cutoff point that kills all other goroutines.
*/

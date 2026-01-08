package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = [3]string{"id1", "id2", "id3"}

var wg = sync.WaitGroup{}

var result = []string{}

// var m = sync.Mutex{}
var m = sync.RWMutex{}

func goroutines() {

	// As we already know Concurrency != Parallelism.
	// Go uses parallelism if the machine has multi-core CPU, which modern machines do.

	fmt.Println("Started")

	for i := range 3 {

		// some_heavy_task(i) // without concurrency (whole program takes 6 secs)

		wg.Add(1)             // increment counter
		go some_heavy_task(i) // with concurrency (just adding `go` tells golang to run it concurrently) (whole program takes 2 secs)
		// Without `WaitGroup`, the program would exit without waiting for the function (which has `go` prefix) to complete.
		// It just assigns a new/idle cpu core, and forgets about it.
		// So, we use `WaitGroup`, simply a counter.

	}

	wg.Wait() // wait for the counter to come back to 0

	fmt.Println("Done:", result)
}

func some_heavy_task(i int) {

	time.Sleep(time.Duration(2000) * time.Millisecond) // placeholder for a task that takes 2 secs

	// println(dbData[i]) // prints in random order, as different cpu cores are assigned
	// (just by a very very small margin, some other core can finish before)

	// If we want to save the data to a slice suppose:
	// result = append(result, dbData[i])
	// Run multiple times and you would notice sometimes we get result like just `[id2, id1]` -> id3 missing.
	// Data is lost as multiple cores insert data at the same time, no mutual exclusion lock is there.
	// To solve this issue, we use mutex lock:
	m.Lock() // while the lock is applied by any cpu core, other cores can't go inside this block of code
	result = append(result, dbData[i])
	m.Unlock() // only once the lock is lifted
	// (This is like semaphores we studied in OS.)

	// Mutex: Only one goroutine can enter `Lock()`.
	// RWMutex: Multiple goroutines can enter `RLock()` (for reading). For writing, we can still use `Lock()`.
	//          When goroutine(s) are inside `RLock()`, any goroutine can't go inside `Lock()`.
	//          In case when one goroutine is on `Lock`, one on `RLock` at the very same time, priority is given to writer.
	// So basically, reading can happen in parallel (as it isn't modifying the data),
	// writing is exclusive, no other writers as-well-as no readers (as if readers are allowed while writing, flawed/old data can be returned).
	// Use-case: When the data is being modified continuously but we want to allow the client to read it meanwhile.
	// Quiz: If we change the loop to run 1000 times, how much time would it take to run completely?
	// If ans. is ~2 secs, why, if not, why not?
	// [For answer, watch the video from 45:58 to 46:41.]
	m.RLock()
	println(dbData[i])
	m.RUnlock()
	// For more explanation, read:
	// https://gemini.google.com/share/e30db4a06a7e
	// https://chatgpt.com/share/6948f260-2150-800a-8457-565c9871806b

	wg.Done() // decrement counter
}

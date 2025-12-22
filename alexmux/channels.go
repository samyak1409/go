package main

import (
	"fmt"
	"math/rand"
	"time"
)

func channels() {

	// Channels: Way to pass around info b/w goroutines.
	// - Hold data
	// - Thread safe (Avoid race conditions)
	// - We can listen for data from channels

	fmt.Println("Channels:")

	// c := make(chan int) // init an (unbuffered (takes only one val of specified type)) channel
	// c <- 7  // add data to channel
	// x := <- c  // access data from channel
	// fmt.Println(x)  // use the data
	// This raises: "fatal error: all goroutines are asleep - deadlock!"
	// Because channels are not supposed to use without goroutines, as you can see, above code makes no sense, we could've directly
	// added the data to `x`.
	// Why deadlock?
	// Because `c <- 7` is executed, and then blocks the subsequent code (in the same function), and waits for something which can pull/use
	// the data from the channel. In above case, that line is below the adding to channel line, hence the deadlock.
	// (And the reason why channels are built to be used with goroutines, as concurrent/parallel execution / non-blocking code.)

	c1 := make(chan int)
	go process1(c1)
	fmt.Println("Chan 1: Data fetched:", <-c1) // access/use the data from channel
	// As soon as `<-c1` is done, `process1` (which is a goroutine i.e. running on a different cpu core) resumes in parallel.
	// Since `channels()` and `process1()` are running in parallel, anyone can print "Chan 1: ..." before the other.

	// What if we want to add multiple vals to the channel and access OTG:
	// We can, even if the unbuffered channel can hold only single val, but when we pull the data/val from the channel, it becomes empty again.
	// So, we can use loop, and add -> use -> add -> use ...
	c2 := make(chan int)
	go process2(c2)
	// If we know the number of times data in the channel would be added, we can use loop with fixed iterations:
	for range 5 {
		fmt.Println(<-c2)
	}

	// Else, we want to keep looping until data is being added in the channel, we can do by using `range c` and `c.close()`:
	c3 := make(chan int)
	go process3(c3)
	// Directly loop on channel, and `val` would be the pulled value:
	for val := range c3 {
		fmt.Println(val)
	}

	// To see a real example of above: 51:15 - 52:51 in the video.

	fmt.Println("Buffered Channels:")
	// Buffered channels (storing multiple values without blocking and waiting to be accessed first):
	// Use case: When the function which is accessing the data from the channel has processing, and so is slow, if using unbuffered channel,
	// that would block the function which is putting data in the channel.
	// So, if we don't want that, we can make channel with a buffer, so that function can add the data to the buffer, which can then be accessed by
	// the primary function later.
	c4 := make(chan int, 4) // (remove the size arg to see "Returning from Process 4..." is printing in the end)
	go process4(c4)
	for val := range c4 {
		fmt.Println(val)
		time.Sleep(time.Second)
	}

}

func process1(ch chan int) {
	ch <- 7 // add data to channel
	fmt.Println("Chan 1: Data added")
}

func process2(ch chan int) {
	for i := range 5 {
		ch <- i
	}
}

func process3(ch chan int) {
	randNum := rand.Intn(10) // calling function (`channels()`) doesn't know how many times data would be added to channel
	fmt.Printf("Chan 3: Will run %d time(s)\n", randNum)
	for i := range randNum {
		ch <- i
	}
	close(ch) // notify the loop in calling function that no more vals are coming, else deadlock error as that for-loop would be waiting.

	// Side-note: We can also use `defer close(ch)` and put it in the beginning, if we want.
}

func process4(ch chan int) {
	for i := range 4 {
		ch <- i
	}
	close(ch)
	fmt.Println("Returning from Process 4...")
}

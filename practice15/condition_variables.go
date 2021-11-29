package practice15

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond // Define global variables

func producer2(out chan<- int, idx int) {
	for {
		// Lock first
		cond.L.Lock()
		// Determine whether the buffer is full
		for len(out) == 3 {
			cond.Wait()
		}
		num := rand.Intn(800)
		out <- num
		fmt.Printf("The first%d Producers, production: %d\n", idx, num)
		// Access to the public area ends, and printing ends, unlock
		cond.L.Unlock()
		// Wake up consumers blocked on condition variables
		cond.Signal()
	}
}

func consumer2(in <-chan int, idx int) {
	for {
		// Lock first
		cond.L.Lock()
		// Determine whether the buffer is empty
		for len(in) == 0 {
			cond.Wait()
		}
		num := <-in
		fmt.Printf("---The first%d Consumers, consumption: %d\n", idx, num)
		// After accessing the public area, unlock
		cond.L.Unlock()
		// Wake up producers blocked on condition variables
		cond.Signal()
	}
}

func Run() {
	// Set random seed number
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int, 3)

	cond.L = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go producer2(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go consumer2(ch, i+1)
	}
	time.Sleep(time.Second * 1)
}

package practice13

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	N = 10
)

func Run() {

	fmt.Println("Goroutine closures pitfall !!!")

	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(id int) {
			time.Sleep(time.Duration(rand.Intn(3111)) * time.Millisecond)
			fmt.Println("Process : ", id)
			wg.Done()
		}(i)
	}

	wg.Wait()

}

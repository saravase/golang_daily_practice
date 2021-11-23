package practice12

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Run() {

	fmt.Println("Multiple routine !!!")

	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup

	producer("optimus1", &wg)
	producer("optimus2", &wg)
	producer("optimus3", &wg)
	producer("optimus4", &wg)

	wg.Wait()

}

func producer(name string, wg *sync.WaitGroup) {

	timeRand := rand.Intn(3)
	timeUp := time.After(time.Duration(timeRand) * time.Millisecond)

	wg.Add(1)

	go func() {

		defer wg.Done()

		for {
			select {
			case <-timeUp:
				fmt.Println(name, " times up")
				return
			default:
				fmt.Println(name)
			}
		}
	}()
}

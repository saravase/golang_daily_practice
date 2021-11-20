package practice9

import (
	"fmt"
	"sync"
)

func Run() {

	fmt.Println("Fanin pattern !!!")

	wg := &sync.WaitGroup{}

	ch1 := producer(wg, 1)
	ch2 := producer(wg, 2)
	ch3 := producer(wg, 3)
	ch4 := producer(wg, 4)

	go consumer(ch1, ch2, ch3, ch4)
	wg.Wait()
}

func producer(wg *sync.WaitGroup, n int) chan int {

	ch := make(chan int)

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i * n
		}
		close(ch)
		wg.Done()
	}()

	return ch
}

func consumer(ch1, ch2, ch3, ch4 <-chan int) {
	for {
		select {
		case n1, ok := <-ch1:
			if ok {
				fmt.Println(n1)
			}
		case n2, ok := <-ch2:
			if ok {
				fmt.Println(n2)
			}
		case n3, ok := <-ch3:
			if ok {
				fmt.Println(n3)
			}
		case n4, ok := <-ch4:
			if ok {
				fmt.Println(n4)
			}
		}
	}
}

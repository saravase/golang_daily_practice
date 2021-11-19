package practice8

import (
	"fmt"
)

func Run() {

	fmt.Println("Generator pattern !!!")

	ch := make(chan bool, 1)

	c := producer(ch)
	go consumer(c)
	<-ch
}

func producer(ch chan bool) chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
		ch <- true
	}()

	return c
}

func consumer(c chan int) {
	for {
		select {
		case n := <-c:
			fmt.Println(n)
		default:
			fmt.Println("Waiting...")
		}
	}
}

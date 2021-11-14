package practice3

import "fmt"

func Run() {

	fmt.Println("Receive only channel")

	ch := make(chan int, 5)

	producer(ch)
	consumer(ch)
}

func consumer(ch <-chan int) {
	for n := range ch {
		fmt.Printf("Value : %v\n", n)
	}
}

func producer(ch chan<- int) {
	ch <- 3
	ch <- 5
	ch <- 6
	close(ch)
}

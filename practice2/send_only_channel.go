package practice2

import "fmt"

// The syntax:
//	chan <- T

func Run() {

	fmt.Println("Send only channel !!!")

	c := make(chan int, 8)

	producer(c)

	for v := range c {
		fmt.Print(v)
	}
}

func producer(c chan<- int) {
	c <- 2
	c <- 6
	c <- 17
	close(c)
}

package practice4

import "fmt"

func Run() {

	fmt.Println("Function returning channel ...")

	ch := producer()
	out := process1(ch)

	for n := range out {
		fmt.Printf("Val := %v\n", n)
	}

}

func producer() <-chan int {

	ch := make(chan int, 10)
	ch <- 3
	ch <- 5
	ch <- 8

	close(ch)
	return ch
}

func process1(in <-chan int) <-chan int {

	out := make(chan int, 10)

	for n := range in {
		out <- (n * 2)
	}

	close(out)
	return out
}

package practice14

import "fmt"

func Run() {
	fmt.Println("Fibanacci Series !!!")

	series := make(chan int)
	quit := make(chan bool)

	n := 10

	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-series)
		}
		quit <- true
	}(n)

	fibonacci(series, quit)

}

func fibonacci(series chan<- int, quit chan bool) {

	x, y := 0, 1

	for {
		select {
		case series <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit...")
			return
		}
	}

}

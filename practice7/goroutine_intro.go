package practice7

import "fmt"

func Run() {

	fmt.Println("Introducing goroutine !!!")

	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i * 2
		}
		close(ch)
	}()

	consumer(ch)
}

func consumer(ch chan int) {
	for num := range ch {
		fmt.Printf("Val :%v\n", num)
	}
}

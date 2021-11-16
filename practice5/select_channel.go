package practice5

import "fmt"

func Run() {
	fmt.Println("Select channel")

	ch := producer()
	for letter := range ch {
		fmt.Print(letter)
	}
	fmt.Println()

}

func producer() <-chan string {
	ch := make(chan string, 24)

	for n := 1; n < 25; n++ {
		select {
		case ch <- "1":
		case ch <- "2":
		case ch <- "3":
		case ch <- "4":
		case ch <- "5":
		case ch <- "6":
		case ch <- "7":
		case ch <- "8":
		case ch <- "9":
		case ch <- "a":
		case ch <- "b":
		case ch <- "c":
		case ch <- "d":
		case ch <- "e":
		case ch <- "f":
		}
	}

	close(ch)
	return ch
}

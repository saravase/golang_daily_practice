package practice6

import (
	"fmt"
	"math/rand"
	"time"
)

func Run() {

	rand.Seed(time.Now().Unix())

	fmt.Println("More chennal selection")

	now := time.Now()

	timeout := time.After(200 * time.Millisecond)

	query := "string"

	ch1 := search(query)
	ch2 := search(query)
	ch3 := search(query)

	select {
	case t := <-ch1:
		fmt.Printf("Search1 completed after : %v\n", t.Sub(now))
	case t := <-ch2:
		fmt.Printf("Search2 completed after : %v\n", t.Sub(now))
	case t := <-ch3:
		fmt.Printf("Search3 completed after : %v\n", t.Sub(now))
	case <-timeout:
		fmt.Println("Too slow, Search timeout after 200ms")
	}
}

func search(query string) <-chan time.Time {
	out := time.After(time.Duration(rand.Intn(2000)) * time.Millisecond)
	return out
}

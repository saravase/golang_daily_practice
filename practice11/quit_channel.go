package practice11

import (
	"fmt"
	"sync"
	"time"
)

func Run() {

	fmt.Println("Quit channel !!!")

	mtx := &sync.Mutex{}

	quit := make(chan bool)

	ch1 := producer("opti1", quit, mtx)
	ch2 := producer("opti2", quit, mtx)
	ch3 := producer("opti3", quit, mtx)
	ch4 := producer("opti4", quit, mtx)
	ch5 := producer("opti5", quit, mtx)

	go func() {

		for {
			select {
			case val := <-ch1:
				fmt.Println(val)
			case val := <-ch2:
				fmt.Println(val)
			case val := <-ch3:
				fmt.Println(val)
			case val := <-ch4:
				fmt.Println(val)
			case val := <-ch5:
				fmt.Println(val)
			}
		}

	}()

	time.Sleep(1 * time.Second)
	quit <- true
	quit <- true
	quit <- true
	quit <- true
	quit <- true
}

func producer(n string, quit chan bool, mtx *sync.Mutex) <-chan string {

	out := make(chan string)

	go func() {
		for {
			select {
			case <-quit:
				mtx.Lock()
				fmt.Println(n, "is quitting")
				mtx.Unlock()
				break
			default:
				out <- n
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	return out

}

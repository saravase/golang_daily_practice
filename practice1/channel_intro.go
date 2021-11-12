package practice1

import "fmt"

func Run() {

	fmt.Println("Chennel introduction !!!")

	// First thing I would like to initalize channel with type int
	var ch chan int
	// Then I want to print the initial value of the chennal
	fmt.Println(ch) // Its return <nil>

	// I will make buffer channel with data type int and buffer size is 2.
	// It means we can send only two values at a time. Then, Those passesd values read from anywhere,
	// we send some other values. So, this is happen continuously.
	ch = make(chan int, 2)

	// we can able to check capacity of the chennal
	fmt.Println(cap(ch)) // 2

	// Put value into this channel
	ch <- 4
	ch <- 8

	// we can able to check length of the chennal
	fmt.Println(len(ch)) // 2

	// Next thing I will try to read value from the nil channel
	var v = <-ch
	fmt.Printf("v : %v\n", v) // Through deadlock in console

	ch <- 18

	v = <-ch
	fmt.Printf("v : %v\n", v)

	ch <- 28

	v = <-ch
	fmt.Printf("v : %v\n", v)

	// Note:
	// When we perform send value into the nil channel, Its is also thrown deadlock

}

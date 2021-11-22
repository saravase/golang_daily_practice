package practice10

import (
	"context"
	"fmt"
	"time"
)

func producer(ctx context.Context, names []string) <-chan string {

	out := make(chan string)

	go func() {
		for _, name := range names {
			select {
			case <-ctx.Done():
				return
			default:
				out <- name
				time.Sleep(time.Second * 2)
			}
		}
		defer close(out)
	}()
	return out
}

func sink(ctx context.Context, in <-chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-in:
			if ok {
				fmt.Println(v)
			}
		}
	}
}

func Run() {

	names := []string{
		"optimus",
		"prime",
		"saravana",
		"kumar",
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()

	ch := producer(ctx, names)
	sink(ctx, ch)

}

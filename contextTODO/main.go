package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processData(val int, ctx context.Context) int {

	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		close(ch)
	}()

	select {
	case <-ch:
		return val*2
	case <-ctx.Done():
		return 0
	}

}



func main() {

	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := range 10 {
			in <- i
		} 
		close(in)
	}()


	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	now := time.Now()

	processParallel(in, out, 5, ctx)

	for v := range out {
		fmt.Println(v)
	}

	fmt.Println(time.Since(now))


}

func processParallel(in <-chan int, out chan<- int, workersNum int, ctx context.Context) {

	wg := sync.WaitGroup{}
	wg.Add(workersNum)

	for range workersNum {
		go func() {
			for v := range in{
				select{
				case <-ctx.Done():
					wg.Done()
					return
				
				case out <- processData(v, ctx):
				}
			}
		wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

}
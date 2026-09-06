package main

import (
	"fmt"
	"sync"
	"time"
)

func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := range 10 {
			ch <- i + 1
		}
		close(ch)
	}()

	return ch
}

func doubler(ch <-chan int) <-chan int {
	ch2 := make(chan int)

	go func() {
		for value := range ch {
			ch2 <- value * 2
			time.Sleep(time.Millisecond * 500)
		}
		close(ch2)
	}()

	return ch2
}

func reader(ch <-chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Println(value)
		}
	}()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	reader(doubler(writer()), &wg)

	wg.Wait()
	fmt.Println("Программа успешно завершена!")
}

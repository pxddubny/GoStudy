package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)


func randomTimeWork() {
	time.Sleep(time.Duration(rand.IntN(5))* time.Second)
}

func predictableTimeWork() {
	ch := make(chan struct{})

	go func() {
		randomTimeWork()
		close(ch)
	}()

	select {
	case <-ch:
		fmt.Println("by func")
	case <-time.After(time.Second * 3):
		fmt.Printf("by after")
	}
}


func main() {

		predictableTimeWork()


}
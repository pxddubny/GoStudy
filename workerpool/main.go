package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type WorkerPool struct {
	jobsCh <-chan string
	doneCh chan<- string
}

func NewWorkerPool(jobsCh <-chan string, doneCh chan<- string) *WorkerPool{
	
	wp := &WorkerPool{
		jobsCh: jobsCh,
		doneCh: doneCh,
	}

	return wp
}

func (wp *WorkerPool) Start(workersNum int) error{

	if workersNum <= 0 {
		return errors.New("stupid?")
	}

	var wg sync.WaitGroup

	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			
			for cite := range wp.jobsCh {
				res, err := Ping(cite)
				if err == nil {
					wp.doneCh <- strconv.Itoa(res) + " Succeed! cite: " + cite
				} else {
					wp.doneCh <- err.Error() + " Failed! cite: " + cite
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(wp.doneCh)
	}()

	return nil
}

func Ping(url string) (int, error) {
	time.Sleep(time.Duration(rand.Intn(201)+50) * time.Millisecond)

	if rand.Intn(5)+1 == 1 {
		return 0, errors.New("connection timeout")
	} else {
		return 200, nil
	}
}

func main() {
	jobsCh := make(chan string, 3)
	doneCh := make(chan string,3)

	wp := NewWorkerPool(jobsCh, doneCh)

	wp.Start(3)

	go func() {
		for i := 0; i < 40; i++ {
			jobsCh <- "cite" + strconv.Itoa(i)
		}
		close(jobsCh)
	}()

	for res := range doneCh {
		fmt.Println(res)
	}
}

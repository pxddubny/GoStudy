package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)


type ConcurrentCache struct {

	mu sync.RWMutex
	data map[string]string

}

func NewConcurrentCache() *ConcurrentCache {

	return &ConcurrentCache{
		data: make(map[string]string),
	}

}

func (cache *ConcurrentCache)Set(key, value string) {

	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.data[key] = value

}

func (cache *ConcurrentCache)Get(key string) (string, bool) {

	cache.mu.RLock()
	defer cache.mu.RUnlock()
	value, ok := cache.data[key]; return value, ok

}

func main() {

	cache := NewConcurrentCache()

	wg := sync.WaitGroup{}
	wg.Add(20)


	for range 10 {

		go func() {
			for i := range "abcdefg" {
				cache.Set(strconv.Itoa(i),"2")
			}
			wg.Done()
		}()

	}

	time.Sleep(time.Second)

	for range 10 {

		go func() {
			for i := range "abcdefg" {
				fmt.Println(cache.Get(strconv.Itoa(i)))
			}
			wg.Done()
		}()

	}

	wg.Wait()

}
/*Реализовать конкурентную запись данных в map.*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	//RW-Mutex
	var mu sync.RWMutex
	m := make(map[int]int)
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = rand.Intn(100)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	for k, v := range m {
		fmt.Println("key:", k, ", val:", v)

	}
	fmt.Printf("\n=====================================================\n")

	// sync.Map
	var w sync.WaitGroup
	var sm sync.Map
	for i := 1; i < 10; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			sm.Store(i, rand.Intn(100))
		}(i)
	}
	w.Wait() // ждем пока горутины закинуть в map
	sm.Range(func(key, value any) bool {
		fmt.Println("key:", key, ", val:", value)
		return true
	})
}

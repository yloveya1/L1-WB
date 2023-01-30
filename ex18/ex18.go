/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	mu *sync.Mutex
	c  uint64
}

func NewCounter() *Counter {
	return &Counter{
		mu: &sync.Mutex{},
	}
}

func (c *Counter) ByMutex() {
	c.mu.Lock()
	c.c++
	c.mu.Unlock()
}

func (c *Counter) GetCount() uint64 {
	return c.c
}

func (c *Counter) ByAtomic() {
	atomic.AddUint64(&c.c, 1)
}

func main() {
	wg := sync.WaitGroup{}
	c1, c2 := NewCounter(), NewCounter()

	for i := 0; i < 50; i++ { // Реализация счетчика с помощью мьютекса для предотвращения гонки данных
		wg.Add(1)
		go c1.ByMutex()
		wg.Done()
	}

	for i := 0; i < 50; i++ { // Реализация счетчика с помощью Atomic
		wg.Add(1)
		go c2.ByAtomic()
		wg.Done()
	}
	wg.Wait()
	fmt.Println("Значение счетичка (Mutex):", c1.GetCount())
	fmt.Println("Значение счетчика (Atomic):", c2.GetCount())
}

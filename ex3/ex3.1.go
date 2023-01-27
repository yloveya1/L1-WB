/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var mass []int = []int{2, 4, 6, 8, 10}

func AtomicSum() {
	var wg sync.WaitGroup
	var sum int64 = 0
	for _, num := range mass {
		wg.Add(1)
		go func(num int) {
			atomic.AddInt64(&sum, int64(num*num))
			wg.Done()
		}(num)
	}
	wg.Wait()
	fmt.Println(sum)
}

func MutexSum() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	var sum1 int
	for _, num := range mass {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			sum1 += num * num
			mu.Unlock()
		}(num)
	}
	wg.Wait()
	fmt.Println(sum1)
}

func ChanelSum() {
	chanout := make(chan int, len(mass))

	sum3 := 0
	wg := sync.WaitGroup{}
	for _, val := range mass {
		wg.Add(1)
		go func(val int) {
			chanout <- val * val
			wg.Done()
		}(val)
	}
	wg.Wait()
	close(chanout)
	for v := range chanout {
		sum3 += v
	}
	fmt.Println(sum3)
}

func main() {
	AtomicSum()
	MutexSum()
	ChanelSum()
}

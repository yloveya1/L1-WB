/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mass []int = []int{1, 2, 3, 4, 5}
	chW, chR := make(chan int), make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		writer(chW, chR)
	}()

	go func() {
		defer wg.Done()
		for _, v := range mass {
			chW <- v
		}
		close(chW)
	}()

	go func() {
		wg.Wait()
		close(chR)
	}()

	for v := range chR {
		fmt.Println(v)
	}
	fmt.Println("Stop working!")
}

func writer(chW, chR chan int) {
	for {
		select {
		case val, ok := <-chW:
			if !ok {
				return
			}
			chR <- val * val
		}
	}
}

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//Method 1: we use channel to signal gorutine that it have to stop.
	quit := make(chan int)
	wg := sync.WaitGroup{}
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Stop gorutine №1\n")
				return
			}

		}
	}()
	quit <- 0

	// 2: Закрытие с помощью контекста
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("Stop gorutine №2\n")
			return
		}
	}()
	cancel()
	wg.Wait()

	// 3: Закрытие по таймеру
	wg.Add(1)
	go func() {
		for {
			select {
			case <-time.After(time.Second / 2):
				fmt.Println("Stop gorutine №3\n")
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()

	// 4: Закрытие канала
	wg.Add(1)
	ch := make(chan bool)
	go func() {
		defer wg.Done()
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					fmt.Println("Stop gorutine №4")
					return
				}
			}
		}

	}()
	close(ch)
	wg.Wait()
}

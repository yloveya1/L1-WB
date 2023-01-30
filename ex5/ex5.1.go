/*Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func writer(ctx context.Context, ch chan int) {
	for {
		select {
		case ch <- rand.Intn(100):
		case <-ctx.Done():
			return
		}
	}
}

func reader(ctx context.Context, ch chan int) {
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	var N float64
	ch := make(chan int)
	defer close(ch)
	var wg sync.WaitGroup
	wg.Add(2)
	for {
		fmt.Printf("Введите количество секунд:")
		_, err := fmt.Scan(&N)
		if err != nil {
			fmt.Println("Некорректное количство секунд. Поробуйте снова")
			continue
		}
		break
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(N)*time.Second) // создание контекста с таймаутом
	defer cancel()
	go func() {
		defer wg.Done()
		reader(ctx, ch) // читающая горутина
	}()
	go func() {
		defer wg.Done()
		writer(ctx, ch) // пишущая горутина
	}()
	wg.Wait()
	fmt.Println("Stop working!")
}

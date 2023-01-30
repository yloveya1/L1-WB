/*Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func writers(timer <-chan time.Time, ch chan int) {
	for {
		select {
		case ch <- rand.Intn(100):
		case <-timer:
			close(ch)
			return
		}
	}
}

func readers(timer <-chan time.Time, ch chan int) {
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(val)
		case <-timer:
			return
		}
	}
}

func main() {
	var N float64
	ch := make(chan int)
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
	timer := time.After(time.Duration(N) * time.Second)
	go func() {
		defer wg.Done()
		readers(timer, ch)
	}()
	go func() {
		defer wg.Done()
		writers(timer, ch)
	}()
	wg.Wait()
	time.Sleep(time.Second)
	fmt.Println("Stop working!")
}

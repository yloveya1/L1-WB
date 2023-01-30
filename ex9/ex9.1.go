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

	var wg sync.WaitGroup // создаем wait-группу, чтобы дождаться завершения всех горутин
	wg.Add(2)
	go func() {
		defer wg.Done()
		writer(chW, chR)
	}()

	go func() {
		defer wg.Done()
		for _, v := range mass { // пишем значения из массива в канал
			chW <- v
		}
		close(chW) // после завершения цикла закрываем канал
	}()

	go func() {
		wg.Wait()
		close(chR) // закрываем канал после того, как отработают все воркеры
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
			if !ok { // если канал закрыт, выходим из функции
				return
			}
			chR <- val * val
		}
	}
}

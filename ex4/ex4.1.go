/*Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора
количества воркеров при старте.Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func writer(ctx context.Context, chanel chan int) {
	for {
		select {
		case chanel <- rand.Intn(100):
		case <-ctx.Done():
			return
		}
	}
}

func reader(ctx context.Context, chanel chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-chanel:
			fmt.Println(val)
		}
	}
}

func main() {
	var count int
	for {
		fmt.Printf("Введите количество воркеров: ")
		_, err := fmt.Scan(&count)
		if err != nil {
			fmt.Println("Некорректное количество воркеров, попробуйте снова")
			continue
		}
		break
	}
	ctx, cancel := context.WithCancel(context.Background())
	chanel := make(chan int)
	for i := 0; i < count; i++ {
		go writer(ctx, chanel)
		go reader(ctx, chanel)
	}
	go func() {
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
		<-signalCh
		cancel()
	}()
	time.Sleep(time.Second)
	fmt.Println("Stop working!")

}

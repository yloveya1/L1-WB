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
		case <-ctx.Done(): // если пришел сигнал завершения при закрытии контекста
			return
		}
	}
}

func reader(ctx context.Context, chanel chan int) {
	for {
		select {
		case <-ctx.Done(): // если пришел сигнал завершения при закрытии контекста
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
		go writer(ctx, chanel) // пишущая горутина
		go reader(ctx, chanel) // читающая горутина
	}
	go func() {
		signalCh := make(chan os.Signal, 1)                    // канал для сигналов системы
		signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM) // ждем сигнал ОС
		<-signalCh
		cancel() // отмена контекста после поступления сигнала ОС
	}()
	time.Sleep(time.Second)
	fmt.Println("Stop working!")
}

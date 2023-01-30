// Реализовать собственную функцию sleep.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	ByContext(time.Second * 3) //
	ByTimeAfter(time.Second * 2)
	fmt.Printf("Время выполнения программы %v", time.Since(t))
}

func ByTimeAfter(t time.Duration) {
	timer := time.After(t) // Time. After ожидает истечения продолжительности t, а затем отправляет текущее время по возвращаемому каналу
	<-timer
}

func ByContext(t time.Duration) {
	ctx, _ := context.WithTimeout(context.Background(), t) // отмена контекста по истечении времени t
	<-ctx.Done()
}

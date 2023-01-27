/* Написать программу, которая конкурентно рассчитает значение квадратов
чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.*/

package main

import (
	"fmt"
	"time"
)

func main() {
	mass := []int{2, 4, 6, 8, 10}
	for _, v := range mass {
		go func(v int) {
			fmt.Printf("%d * %d = %d\n", v, v, v*v)
		}(v)
	}
	time.Sleep(1 * time.Second)
}

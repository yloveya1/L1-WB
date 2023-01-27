/* Написать программу, которая конкурентно рассчитает значение квадратов
чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout. */

package main

import (
	"fmt"
	"sync"
)

func main() {
	mass := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(len(mass))
	for _, v := range mass {
		go func(v int) {
			defer wg.Done()
			fmt.Printf("%d * %d = %d\n", v, v, v*v)
		}(v)
	}
	wg.Wait()
}

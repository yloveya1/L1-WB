// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
)

func main() {
	arr := []int{-5, -1, 1, 2, 4, 9, 25} // в качестве исходных данных - отсортированный массив
	findDigit := 27
	result := binarySearch(arr, findDigit)
	if result < 0 {
		fmt.Printf("Число %d отсутствует в слайсе", findDigit)
	} else {
		fmt.Printf("arr[%d] = %d", result, findDigit)
	}
}

func binarySearch(ar []int, n int) int {
	low, high := 0, len(ar)-1
	for low <= high {
		mid := low + (high-low)/2
		if n < ar[mid] { // значение искомого числа меньше среднего => первая половина массива
			high = mid - 1 // меняем значение конечного индекса
		} else if n > ar[mid] { // значение искомого числа больше среднего => вторая половина массива
			low = mid + 1 // меняем значение начального индекса
		} else {
			return mid // возвращаем полученный индекс
		}
	}
	return -1 // возвращаем -1 в случае ошибки
}

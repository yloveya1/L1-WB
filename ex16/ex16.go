/*Реализовать быструю сортировку массива (quicksort) встроенными методами языка.*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	ar := []int{-9, 0, -15, 5, 2, 9, 8, 0}
	QuickSort(ar, 0, len(ar)-1)
	fmt.Println(ar)
	fmt.Println("============================")
	ar = []int{-9, 0, -15, 5, 2, 9, 8, 0}
	sort.Ints(ar)
	fmt.Println(ar)
}

func QuickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = QuickSort(arr, low, p-1)
		arr = QuickSort(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

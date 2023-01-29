// Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Printf("a: %v, b: %v\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a: %v, b: %v\n", a, b)

	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

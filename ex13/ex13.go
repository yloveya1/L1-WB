// Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main() {
	// Case 1
	a, b := 1, 2
	a, b = b, a
	fmt.Printf("a: %v, b: %v\n", a, b)
	// Case 2
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a: %v, b: %v\n", a, b)
	// Case 3 by XOR
	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("a: %v, b: %v", a, b)
}

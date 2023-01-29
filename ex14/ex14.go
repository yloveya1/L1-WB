/*Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.*/

package main

import (
	"fmt"
	"reflect"
)

func getType1(temp interface{}) {
	switch val := temp.(type) {
	case int:
		fmt.Println("int: ", val)
	case string:
		fmt.Println("string: ", val)
	case bool:
		fmt.Println("bool: ", val)
	case chan interface{}:
		fmt.Println("chan interface{}: ", val)
	}
}

func getType2(temp interface{}) {
	fmt.Println(reflect.TypeOf(temp))
}

func main() {
	integer := 4
	str := "Hello"
	boolean := true
	channel := make(chan interface{})

	getType1(integer)
	getType1(str)
	getType1(boolean)
	getType1(channel)

	fmt.Println("================================")

	getType2(integer)
	getType2(str)
	getType2(boolean)
	getType2(channel)
}

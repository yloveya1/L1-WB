/*Реализовать пересечение двух неупорядоченных множеств*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	mass1 := []int{1, 2, 3, 4, 5}
	mass2 := []int{5, 4, 3, 1, 9, 3}
	res := make(map[int]struct{})

	for _, val := range mass1 {
		res[val] = struct{}{}
	}
	for _, val := range mass2 {
		if _, ok := res[val]; ok != false {
			fmt.Println(val)
		}
	}
}

var a = []int{1, 3, 5, 10, 4, 8}
var b = []int{2, 6, 3, 4, 11, 7, 4, 4}

func simpleIntersect(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	av := reflect.ValueOf(a)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		if contains(b, el) {
			set = append(set, el)
		}
	}
	return set
}

func hashIntersect(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	hash := make(map[interface{}]bool)
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		hash[el] = true
	}

	for i := 0; i < bv.Len(); i++ {
		el := bv.Index(i).Interface()
		if _, found := hash[el]; found {
			set = append(set, el)
		}
	}

	return set
}

func contains(a interface{}, e interface{}) bool {
	v := reflect.ValueOf(a)

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

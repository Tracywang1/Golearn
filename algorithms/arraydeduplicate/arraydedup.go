package main

import "fmt"

func arrayDeduplicate(arry []int) []int {
	res := []int{}
	for _, v := range arry {
		if !isInAarray(res, v) {
			res = append(res, v)
		}
	}
	return res
}

func isInAarray(array []int, num int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == num {
			return true
		}
	}
	return false
}

func arrayIntersection(array1 []int, array2 []int) []int {
	res := []int{}
	array1 = arrayDeduplicate(array1)
	array2 = arrayDeduplicate(array2)
	fmt.Println(array1, array2)
	for _, v1 := range array1 {
		if isInAarray(array2, v1) {
			res = append(res, v1)
		}
	}
	return res
}

func main() {
	s := []int{1, 2, 3, 4, 5, 2, 2, 3, 4, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5}
	s1 := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 10, 11, 12, 13}
	// fmt.Println(arrayDeduplicate(s))
	// fmt.Println(arrayDeduplicate(s1))
	fmt.Println(arrayIntersection(s, s1))
}

package main

import (
	"fmt"
)

func twosum(num []int, target int) []int {
	res := []int{}
	for i := 0; i < len(num)-1; i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i]+num[j] == target {
				res = append(res, num[i], num[j])
			}
		}
	}
	return res
}
func main() {
	s := []int{9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 10}
	target := 10
	fmt.Println(twosum(s, target))
}

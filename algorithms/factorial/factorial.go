package main

import (
	"fmt"
)

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
func main() {
	var i uint64 = 15
	fmt.Printf("%v的阶乘是%v\n", i, Factorial(i))
}

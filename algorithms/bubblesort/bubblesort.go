package bubblesort

import (
	"fmt"
)

func BubbleSort(values []int) {
	flag := true
	for i := 0; i < len(values)-1; i++ {
		flag = true
		fmt.Println("values:", i, values[i], len(values))
		for j := 0; j <= len(values)-i-2; j++ {
			fmt.Println("j:", j, values[j])
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}

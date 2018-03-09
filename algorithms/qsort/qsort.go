package qsort

import (
	"fmt"
)

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	fmt.Println("i,j,p,value[left]", i, j, p, values[left])
	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	fmt.Println(values, len(values))
	quickSort(values, 0, len(values)-1)
}

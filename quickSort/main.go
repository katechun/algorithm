package main

import (
	"fmt"
)

func quicksort(array []int) []int {
	var less []int
	var greater []int
	var pivot int
	if len(array) < 2 {
		return array
	} else {

		pivot = array[0]
		for i := 1; i < len(array); i++ {
			if array[i] <= pivot {
				less = append(less, array[i])
			}
		}

		for i := 1; i < len(array); i++ {
			if array[i] > pivot {
				greater = append(greater, array[i])
			}
		}
	}

	return append(append(quicksort(less), pivot), quicksort(greater)...)
}

func main() {
	arr := []int{
		10, 5, 2, 3,
	}
	fmt.Println(quicksort(arr))
}

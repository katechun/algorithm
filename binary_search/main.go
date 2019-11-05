package main

import (
	"fmt"
)

func binary_search(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		guess = guess
		if guess == item {
			return mid
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	var my_list []int
	my_list = []int{
		1, 3, 5, 7, 9,
	}
	gg := binary_search(my_list, 5)
	fmt.Println(gg)
}

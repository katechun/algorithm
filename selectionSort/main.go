package main

import (
	"fmt"
)

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}

	return smallest_index
}

func selectionSort(arr []int) []int {
	newArr := []int{}
	oldArr := arr

	for i := 0; i < len(arr); i++ {
		smallest := findSmallest(oldArr)
		newArr = append(newArr, arr[smallest])

		if smallest != 0 && smallest != len(oldArr)-1 {
			oldArr = append(oldArr[0:smallest], oldArr[smallest+1:]...)
		}

		if smallest == 0 {
			oldArr = oldArr[1:]
			fmt.Println("0:", oldArr)
		}

		if smallest == len(arr)-1 {
			oldArr = oldArr[0 : len(arr)-2]
		}
	}

	return newArr
}

func main() {
	arr := []int{
		5, 3, 6, 2, 10,
	}
	fmt.Println(selectionSort(arr))
}

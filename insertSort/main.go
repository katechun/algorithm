package main

import (
	"fmt"
)

func main(){
	list :=[]int{5,2,4,6,1,3}

	for  j:=1;j<len(list);j++ {
		key := list[j]
		i := j - 1
		for i >= 0 && list[i] > key {
			list[i+1] = list[i]
			i = i - 1
		}
		list[i+1] = key

	}

	fmt.Println(list)
}


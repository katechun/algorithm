package main

import (
	"algorithm/breadthFirstSearch/queue"
	"fmt"
)

func search(name string) bool {
	search_queue := queue.Queue{}
	search_queue.Put(name)
	searched := []string{}

	for search_queue.Size() != 0 {
		person := search_queue.Pop()
		if personExist(person, searched) {
			if person_is_seller(person) {
				fmt.Println(person, "is a mango seller!")
				return true
			}
		}

	}
	return false
}

func person_is_seller(name string) bool {
	return name[:1] == "m"
}

func personExist(person string, searched []string) bool {
	for _, v := range searched {
		if v == person {
			return false
		}
	}
	return true
}

func main() {
	if search("bill") {
		fmt.Println("找到芒果供应商！")
	}
}

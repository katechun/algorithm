package main

import (
	"algorithm/stack/comm"
	"fmt"
)

func main() {
	var stack comm.Stack
	stack.Numbers = []interface{}{1, 3, 5, 7, 9}
	for i := 0; i < stack.Length(); i++ {
		fmt.Println(stack.Numbers[i])
	}
}

package main

import (
	"algorithm/bag/comm"
	"fmt"
	"math"
)

func main() {
	var nums comm.Bag
	var sum int
	var sum2 int
	nums.Numbers = []int{3, 3, 0, 6}
	N := nums.Size(nums)
	for i := 0; i < len(nums.Numbers); i++ {
		sum = sum + nums.Numbers[i]
	}

	mean := sum / N

	for i := 0; i < len(nums.Numbers); i++ {
		sum2 = sum2 + nums.Numbers[i]
	}

	std := math.Sqrt(float64(sum2 / (N - 1)))

	fmt.Println("Mean: ", mean)
	fmt.Println("Std dev: ", std)

}

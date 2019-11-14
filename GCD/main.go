package main

import (
	"fmt"
)

func gcd(p int, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return gcd(q, r)
}

func main() {
	res := gcd(4, 6)
	fmt.Println(res)
}

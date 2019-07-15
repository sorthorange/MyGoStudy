package main

import "fmt"

func Fb(n int) []int {
	x := make([]int, n)
	x[0], x[1] = 1, 1
	for i := 2; i < n; i++ {
		x[i] = x[i-1] + x[i-2]
	}
	return x
}

func main() {
	fmt.Println("", Fb(2))
}

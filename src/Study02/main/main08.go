package main

import "fmt"

func Map(f func(int) int, a []int) []int {
	j := make([]int, len(a))
	for k, v := range a {
		j[k] = f(v)
	}
	return j
}

func main() {
	m := []int{1, 3, 4}
	f := func(i int) int {
		return i * i
	}
	fmt.Printf("%v", (Map(f, m)))
}

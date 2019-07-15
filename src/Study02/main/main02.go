package main

import "fmt"

func sort(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func main() {
	a, b := 7, 2
	a, b = sort(a, b)
	fmt.Printf("%d,%d", a, b)
}

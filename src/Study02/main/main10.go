package main

import "fmt"

func bubbleSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func main() {
	a := []int{9, 5, 1, 3, 4, 8}
	fmt.Printf("%v", bubbleSort(a))
}

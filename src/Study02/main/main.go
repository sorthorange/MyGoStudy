package main

import "fmt"

func Avg(a []float64) (avg float64) {
	var sum float64 = 0
	switch len(a) {
	case 0:
		avg = 0
	default:
		for _, v := range a {
			sum += v
		}
		avg = sum / float64(len(a))
	}
	return
}

func main() {
	a := []float64{1.2, 3.6, 4.4, 5.9}
	fmt.Printf("%.2f\n", Avg(a[0:2]))
}

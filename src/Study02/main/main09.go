package main

func Min(a []int) int {
	var min = a[0]
	for _, t := range a {
		if min > t {
			min = t
		}
	}
	return min
}
func Max(a []int) int {
	var max = a[0]
	for _, t := range a {
		if max < t {
			max = t
		}
	}
	return max
}
func main() {
	a := []int{1, 2, 4, 5, 9}
	print(Min(a))
}

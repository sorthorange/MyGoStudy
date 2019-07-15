package main

import "fmt"

type Sorter interface {
	Len() int           //长度方法
	Less(i, j int) bool //比较方法
	Swap(i, j int)      //swap方法
}

type Xi []int
type Xs []string

func (p Xi) Len() int {
	return len(p)
}
func (p Xi) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Xi) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Xs) Len() int {
	return len(p)
}
func (p Xs) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Xs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(j, i) {
				x.Swap(i, j)
			}
		}
	}
}

func main() {
	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}

	Sort(ints)
	Sort(strings)
	fmt.Printf("%v\n", ints)
	fmt.Printf("%v\n", strings)
}

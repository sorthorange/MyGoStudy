package main

import "sort"

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

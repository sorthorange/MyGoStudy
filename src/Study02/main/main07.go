package main

import "fmt"

func Map02(f func(string) string, a []string) []string {
	j := make([]string, len(a))
	for k, v := range a {
		j[k] = f(v)
	}
	return j
}

func main() {
	m := []string{"123", "456", "159"}
	f := func(s string) string {
		return s
	}
	fmt.Printf("%v", (Map02(f, m)))
}

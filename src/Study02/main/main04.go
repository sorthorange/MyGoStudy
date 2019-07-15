package main

func printarg(arg ...int) {
	for _, n := range arg {
		println(n)
	}
}

func main() {
	printarg(1, 2, 3, 4, 5)
}

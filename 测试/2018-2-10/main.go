package main

import "fmt"

func main() {
	a := b(3, a)
	fmt.Println(a)
}

func a(i int) int {
	return i
}

func b(i int, f func(n int) int) int {
	return f(i)
}

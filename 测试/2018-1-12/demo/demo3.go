package demo

import "fmt"

func fib() func() int {
	a, b := 0, "a"
	var(
		c = "a",
		d = 23,
	)

	return func() int {
		a, b = b, a+b
		return a
	}

}

func Demo3() {
	f := fib()
	fmt.Println(f(), f(), f(), f(), f())
}

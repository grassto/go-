package demo

import "fmt"

type Number *Number

func zero() *Number {
	return nil
}

func isZero(x *Number) bool {
	return x == nil
}

func add1(x *Number) *Number {
	e := new(Number)
	*e = x
	return e
}

func sub1(x *Number) *Number {
	return *x
}

func add(x, y *Number) *Number {
	if isZero(y) {
		return x
	}
	return add(add1(x), sub1(y))
}

func mul(x, y *Number) *Number {
	if isZero(x) || isZero(y) {
		return zero()
	}
	return add(mul(x, sub1(y)), x)
}

func fact(n *Number) *Number {
	if isZero(n) {
		return add1(zero())
	}
	return mul(fact(sub1(n)), n)
}

// -------------------------------------
// Helpers to generate/count Peano integers

func gen(n int) *Number {
	if n > 0 {
		return add1(gen(n - 1))
	}
	return zero()
}

func count(x *Number) int {
	if isZero(x) {
		return 0
	}
	return count(sub1(x)) + 1
}

func Demo4() {
	for i := 0; i <= 9; i++ {
		f := count(fact(gen(i)))
		fmt.Println(i, "! =", f)
	}
}

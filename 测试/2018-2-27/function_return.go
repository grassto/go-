package main

import "fmt"

func example4() {
	p2 := Add2()
	fmt.Println("和为：", p2(3, 3))
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
	fmt.Printf("再加一个3，得到结果为：%v\n", TwoAdder(3))
}

func Add2() func(...int) int {
	return func(b ...int) int {
		sum := 0
		for _, val := range b {
			sum += val
		}
		return sum
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		a += b
		return a + b
	}
}

func example5() {
	f := Adder2()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

// 使用闭包函数可以继续操作外部函数中的局部变量
func Adder2() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

// 用defer改写一下Adder2
func Adder2Def() (f func(int) int) {
	var x int
	defer func() {
		f = func(delta int) int {
			x += delta
			return x
		}
	}()
	return f
}

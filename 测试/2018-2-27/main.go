package main

import "fmt"

func main() {
	// example1()
	// example2()
	// practise6_8()
	// fmt.Println(example3())
	// example4()
	// example5()
	// fmt.Println(example6())
	// example6()
	// example7()
	// example8()
	// example10()
	timeSub()
}

// 直接调用匿名函数，在最后加()
func example1() {
	func() {
		sum := 0
		for i := 1; i < 1e6; i++ {
			sum += i
		}
		fmt.Println("总和：", sum)
	}()
}

// 匿名函数可分配不同的地址
func example2() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		// g的地址都是一样的（是go的版本问题？）
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

func practise6_8() {
	fv := func() {
		fmt.Println("Hello Word")
	}
	fv()
	fmt.Printf("变量fv的类型为：%T", fv)
}

// 返回值为2
func example3() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

// 暂时搁浅
func example6() int {
	var g int
	func(i int) {
		s := 0
		for j := 0; j < i; j++ {
			s += j
		}
		g = s
	}(1000) // Passes argument 1000 to the function literal.
	return g
}

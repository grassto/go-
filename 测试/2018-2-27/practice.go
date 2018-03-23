package main

import "fmt"

// 使用闭包实现斐波那契数列
func fibonacci() func() {
	// x第一个，y第二个，z和，i位置
	var x, y, z, i int
	y = 1
	i = 1
	return func() {
		z = x + y
		x = y
		y = z
		fmt.Printf("fibonacci(%d) is: %d\n", i, z)
		i++
	}
}

func example7() {
	fv := fibonacci()
	for i := 0; i < 200; i++ {
		fv()
	}
}

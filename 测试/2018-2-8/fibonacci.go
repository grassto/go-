package main

import "fmt"

// 从0开始，计算出第几位上的斐波那契数
func fibonacci(a int) (result int) {
	if a <= 1 {
		result = 1
	} else {
		result = fibonacci(a-1) + fibonacci(a-2)
	}
	return
}

func getFibonacci(a int) {
	fmt.Println(fibonacci(a))
}

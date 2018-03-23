package main

const LIM = 41

// 定义全局变量，用来存储计算过的斐波那契数
var fibs []uint64

func fibonacciByMemorization(n uint64) (res uint64) {
	if n < 2 {
		res = 1
	} else {
		res = fibonacciByMemorization(n-2) + fibonacciByMemorization(n-1)
	}
	return
}

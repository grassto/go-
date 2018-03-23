package main

import (
	"fmt"
	"io"
	"log"
)

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

// 用来追踪函数的执行
func defer1() {
	defer un(trace("defer1"))
	fmt.Println("in: defer1")
}

// 打印函数的参数和返回值
func defer2(a string, b int) (n int, err error) {
	defer func() {
		log.Printf("defer2(%q,%d) = %d,%v", a, b, n, err)
	}()
	return 3, io.EOF
}

func defer3() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Println("循环结束")
}

func defer4() {
	for i := 0; i < 5; i++ {
		defer func(n int) { fmt.Printf("%v\n", n) }(i)

		fmt.Println(i)
	}
	fmt.Println("循环结束")
}

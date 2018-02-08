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

func defer1() {
	defer un(trace("defer1"))
	fmt.Println("in: defer1")
}

func defer2(a string, b int) (n int, err error) {
	defer func() {
		log.Printf("defer2(%q,%d) = %d,%v", a, b, n, err)
	}()
	return 3, io.EOF
}

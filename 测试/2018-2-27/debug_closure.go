package main

import (
	"log"
	"runtime"
)

// 使用闭包实现调试，显示文件的位置和行号
func example9() {
	where := func() {
		// 文件位置和行号
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s %d", file, line)
	}

	a := "dds"
	a += "s"
	where()

	ddd := 3
	ddd += 3
	where()
}

func example10() {
	where := log.Print
	a := "dds"
	a += "s"
	where()

	ddd := 3
	ddd += 3
	where()
}

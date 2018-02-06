package main

import (
	"fmt"
	"os"
	"runtime"

	"./random"
	// "./exercise"
)

func main() {
	random.RunRandom()
	// getPath()
	// exercise.RunLocal()
}

func test1() {
	var (
		HOME   = os.Getenv("HOME")
		USER   = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)

	fmt.Println("HOME的值：", HOME, ",USER的值：", USER, ",GOROOT的值：", GOROOT)
}

func getPath() string {
	var goos string = runtime.GOOS
	fmt.Printf("当前的操作系统为：%s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("环境变量为：%s\n", path)
	return goos
}

func testComplex() {
	var c1 complex128 = 4 + 3i
	fmt.Printf("%v\n", c1)
	// 输出实部与虚部
	fmt.Println(real(c1), imag(c1))
}

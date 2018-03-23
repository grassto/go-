package main

import (
	"fmt"
	"strings"
)

// 工厂函数
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasPrefix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func example8() {
	addBmp := MakeAddSuffix(".bmp")
	addJpep := MakeAddSuffix(".jpeg")
	fmt.Println(addBmp("file"))
	fmt.Println(addJpep("file"))
}

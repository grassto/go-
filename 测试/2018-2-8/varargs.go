// 练习6.3
package main

import "fmt"

func practice1(a ...string) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func e1(a ...string) {
	e2(a...)
	e3(a)
}

func e2(a ...string) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func e3(a []string) {
	for _, v := range a {
		fmt.Println(v)
	}
}

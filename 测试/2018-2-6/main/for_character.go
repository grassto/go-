// 练习5.5
package main

import "fmt"

func for3() {
	for i, str := 0, "G"; i < 25; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(str)
		}
		fmt.Println()
	}
}

func for4() {
	for i, str := 25, "G"; i > 0; i-- {
		fmt.Println(str)
		str += "G"
	}
}

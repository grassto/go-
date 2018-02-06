//练习5.4
package main

import "fmt"

func for1() {
	for i := 0; i < 15; i++ {
		fmt.Println(i)
	}
}

func for2() {
	i := 0
FOR2:
	if i < 15 {
		fmt.Println(i)
		i++
		goto FOR2
	}
}

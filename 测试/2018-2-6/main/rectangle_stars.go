// 练习5.8
package main

import "fmt"

func rec() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			switch i {
			case 0, 9:
				fmt.Print("*")
			default:
				switch j {
				case 0, 19:
					fmt.Print("*")
				default:
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}

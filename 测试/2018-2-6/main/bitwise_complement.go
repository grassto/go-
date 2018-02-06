// 练习5.6
package main

import "fmt"

// 使用按位补码从 0 到 10，使用位表达式 %b 来格式化输出。
func for5() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("the complement of %b is: %b\n", i, ^i)
	}
}

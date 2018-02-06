package main

import (
	"fmt"
)

func main() {
	// switchThrough()

	// ss, err := Season(3)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Print(ss)

	// for1()
	// for2()

	// for3()
	// for4()

	// for5()
	rec()
}

func switchThrough() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

// Season 判断输入的月份是什么季节
func Season(month int) (ss string, err error) {
	switch month {
	case 12, 1, 2:
		err = nil
		ss = "冬天"
		// return "冬天", err
	case 3, 4, 5:
		err = nil
		ss = "春天"
		// return "春天", err
	case 6, 7, 8:
		err = nil
		ss = "夏天"
		// return "夏天", err
	case 9, 10, 11:
		err = nil
		ss = "秋天"
		// return "秋天", err
	default:
		ss = ""
		// return nil, err
	}
	return
}

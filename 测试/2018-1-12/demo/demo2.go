package demo

import (
	"flag"
	"fmt"
)

// 返回一个bool类型的指针
var nFlag = flag.Bool("n", true, "末尾是否换行")

func Demo2() {
	flag.Parse()
	s := ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += ""
		}
		s += flag.Arg(i)
	}
	if *nFlag {
		s += "\n"
	}
	fmt.Println(flag.Args())
	// os.Stdout.WriteString(s)
	fmt.Println(s)
}

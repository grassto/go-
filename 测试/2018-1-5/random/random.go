package random

import (
	"fmt"
	"math/rand"
	"time"
)

func RunRandom() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d /", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8) //返回的是[0,8)的整数
		fmt.Printf("%d / ", r)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond()) //种子
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
}

package main

import (
	"fmt"
	"time"
)

func timeSub() {
	start := time.Now()
	example7()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("ongCalculation took this amount of time: %s\n", delta)
}

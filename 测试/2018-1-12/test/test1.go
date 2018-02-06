package test

import "fmt"

func Sumqq(numbers []int) int {
	// 	sum := 0
	//     for index, n := range numbers {
	// 		fmt.Println(index)
	//         sum += n
	//     }

	//     return sum

	sum := 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Println(sum)
	return sum
}

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

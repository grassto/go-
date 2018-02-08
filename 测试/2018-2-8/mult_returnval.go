// 练习6.1
package main

func spd1(num1 int, num2 int) (int, int, int) {
	sum := num1 + num2
	product := num1 * num2
	diff := num1 - num2
	return sum, product, diff
}

func spd2(num1, num2 int) (sum int, product int, diff int) {
	sum, product, diff = num1+num2, num1*num2, num1-num2
	return
}

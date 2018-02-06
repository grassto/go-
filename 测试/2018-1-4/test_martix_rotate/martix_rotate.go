package martix

import "fmt"

// type的使用，类名等价定义
// 在此，g2048和[4][4]int类型等价
type g2048 [4][4]int

// 可以看作是给类添加方法，可以用g2048来点Print调用
func (t *g2048) Print() {
	for _, line := range t {
		for _, number := range line {
			fmt.Printf("%v ", number)
		}
		fmt.Println()
	}
}

// 竖直方向翻转
func (t *g2048) MirrorV() {
	tn := new(g2048)
	for i, line := range t {
		for j, num := range line {
			tn[len(t)-i-1][j] = num
		}
	}
	*t = *tn
}

func TestRotate() {
	fmt.Println("初始数组")
	t := g2048{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 1}, {6, 98}}
	t.Print()
	fmt.Println()

	// 二维数组的长度是以第第一个数组的长度来算的
	fmt.Println("二维数组的长度为：", len(t))
	fmt.Println()

}

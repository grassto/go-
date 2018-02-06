官网的七个例子

1.Hello Word

2.flag处理命令行参数
> go run main.go -n=false good morning
good morning
> go run main.go -n=true good morning
good moring
 
> go run main.go good morning
good moring

3.闭包实现斐波那契

4.阶乘运算
例子中反映了：阶乘依赖于自然数乘法；自然数乘法依赖于自然数加法，而自然数加法其实是个递归定义。
自然数就是一个链表，链表的第一个元素代表 0 ，第二个元素代表 1，第三个 代表 2，一直下去。10 就是一个长度为 10 的链表！求 10！的思路不再是循环一下，乘一下！而是，构造出一个长度是 10！ 的链表，然后 count 一下长度，就是 10！多此一举？！这里所有的操作只建立在“加一”、“减一”和“递归”的基础之上。

5.并发求圆周率

6.并发通过筛选法求素数

7.孔明棋

8.二叉树
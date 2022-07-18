package main

import "fmt"

/**
 * 在前面的例子中， 我们讲过 for 和 range 为基本的数据结构提供了迭代的功能。
 * 我们也可以使用这个语法来遍历的从通道中取值。
 **/
func main() {

	// 我们将遍历在 queue 通道中的两个值。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	/**
	 * range 迭代从 queue 中得到每个值。
	 * 因为我们在前面 close 了这个通道，所以，这个迭代会在接收完 2 个值之后结束。
	 **/
	for elem := range queue {
		fmt.Println(elem)
	}
}

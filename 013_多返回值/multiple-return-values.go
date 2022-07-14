package main

import "fmt"

//(int, int) 在这个函数中标志着这个函数返回 2 个 int。
func valS() (int, int) {
	return 3, 7
}

func main() {

	//这里我们通过 多赋值 操作来使用这两个不同的返回值。
	a, b := valS()
	fmt.Println(a)
	fmt.Println(b)

	//如果你仅仅需要返回值的一部分的话，你可以使用空白标识符 _。
	_, c := valS()
	fmt.Println(c)
}

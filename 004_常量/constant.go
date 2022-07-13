package main

import (
	"fmt"
	"math"
)

// const 用于声明一个常量。
const s string = "constant"

func main() {

	fmt.Println(s)

	// const 语句可以出现在任何 var 语句可以出现的地方
	const n = 50000000

	// 常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)

	// 数值行常量没有确定的类型，直到被给定某个类型，比如显式类型转换
	fmt.Println(int64(d))

	/*	一个数字可以根据上下文的需要（比如变量赋值、函数调用）自动确定类型。
		举个例子，这里的 math.Sin 函数需要一个 float64 的参数，n 会自动确定类型。*/
	fmt.Println(math.Sin(n))

	// 我们尝试在函数内部再定义一个数值打印
	const s int = 1
	fmt.Println(s)
	// 函数内部局部常量与全局不会冲突

	//使用函数验证
	funcConst()
	sConst()
}

func funcConst() {
	const s string = "one"
	fmt.Println(s)
}

func sConst() {
	fmt.Println(s)
}

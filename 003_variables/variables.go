package main

import "fmt"

func main() {

	// var 声明 1 个或者多个变量
	var a = "initial"
	fmt.Println(a)

	// 你可以一次性声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go会自动推断出已经有初始值的变量类型
	var d = true
	fmt.Println(d)

	/**
	* 声明后却没有给出对应的初始值时，变量将会初始化零值。
	* 例如 int 的零值是 0
	* 例如 string 的零值是 ""
	* 例如 map 的零值是 map[]
	* 例如 struct 的零值是 {}
	* 例如 interface 的零值是 nil
	 */

	var e interface{}
	fmt.Println(e)

	// := 语法是声明初始化变量的简写，例如 var f string = "short" 可以简写为下面这样
	f := "short"
	fmt.Println(f)
}

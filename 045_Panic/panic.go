package main

import (
	"fmt"
	"os"
)

/*panic 意味着有些出乎意料的错误发生。
通常我们用它来表示程序正常运行中不应该出现的错误， 或者我们不准备优雅处理的错误。*/
func main() {

	/*我们将使用 panic 来检查这个站点上预期之外的错误。
	  而该站点上只有一个程序：触发 panic。*/

	//panic("a problem")
	/**
	 * panic 的一种常见用法是：当函数返回我们不知道如何处理（或不想处理）的错误值时，中止操作。
	 * 如果创建新文件时遇到意外错误该如何处理？这里有一个很好的 panic 示例。
	 */
	_, err := os.Create("./tmp/file")
	if err != nil {
		panic(err)
	}

	fmt.Println("成功")
}

//运行程序将会导致 panic： 输出一个错误消息和协程追踪信息，并以非零的状态退出程序。

/*当 main 中触发第一个 panic 时，程序就会退出而不会执行代码的其余部分。
如果你想看到程序尝试创建 /tmp/file 文件，请注释掉第一个panic。*/

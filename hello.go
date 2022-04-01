package hello

import "fmt"

// SayHi 一个打招呼的函数
func SayHi(name string, age int) {
	fmt.Printf("你好%s，我是帅哥。今年%d岁，很高兴认识你~\n", name, age)
}

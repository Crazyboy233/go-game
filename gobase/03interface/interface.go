package main

import "fmt"

// 1. interface（接口）
type Speaker interface {
	Speak()
}

type Dog struct{}

func (dog Dog) Speak() {
	fmt.Println("汪汪汪")
}

// 2. 错误处理
func Divid(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return a / b, nil
}

// 3. 闭包：闭包是返回函数的函数，可以捕获外部变量
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	dog := Dog{}
	dog.Speak()

	res, error := Divid(2, 0)
	if error == nil {
		fmt.Println("结果为：", res)
	} else {
		fmt.Println(error)
	}

	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}

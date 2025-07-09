package command

import "fmt"

// 菜单
func BeforeLogin() {
	fmt.Println("欢迎来到游戏世界！")
	fmt.Println("请选择操作:")
	fmt.Println("1.登录")
	fmt.Println("2.注册")
	fmt.Println("0.退出")
}

func AfterLogin() {
	fmt.Println("请选择你的操作")
	fmt.Println("1. 匹配")
	fmt.Println("2. 我的")
	fmt.Println("3. 登出")
	fmt.Println("4. 退出")
}

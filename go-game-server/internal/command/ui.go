package command

import "fmt"

// 登录之前的菜单
func BeforeLogin() {
	fmt.Println("===============欢迎来到游戏世界！===============")
	fmt.Println("请选择操作:")
	fmt.Println("1.登录")
	fmt.Println("2.注册")
	fmt.Println("0.退出")
}

// 登录之后的菜单
func AfterLogin() {
	fmt.Println("===============欢迎来到游戏世界！===============")
	fmt.Println("请选择你的操作")
	fmt.Println("1. 匹配")
	fmt.Println("2. 我的")
	fmt.Println("3. 登出")
	fmt.Println("4. 退出")
}

// 开发者模式菜单
func Debug() {
	fmt.Println("===============你已进入开发者模式===============")
	fmt.Println("请输入命令：")
	// fmt.Println("-1: 退出开发者模式")
}

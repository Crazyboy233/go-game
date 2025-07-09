package main

import (
	"fmt"
	"go-game-server/internal/command"
	"go-game-server/internal/player"
)

const (
	StateInit = iota
	StateLoggedIn
)

var currentUser *player.Player = nil

func main() {
	for {
		// 登录前的菜单
		command.BeforeLogin()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// 执行登录操作
			fmt.Print("请输入你的账号:")
			// 构造 Command 绕过 paraer 直接走dispatcher分发逻辑。
			var id string
			fmt.Scan(&id)

			input := command.Command{Name: "/Login", Args: []string{id}}
			ok := command.Dispatch(input)
			if ok {
				for {
					command.AfterLogin()

					var choice string
					fmt.Scan(&choice)
				}
			} else {
				continue
			}

		case 2:
			// 执行注册操作
			fmt.Print("请输入你的名称：")
			var name string
			fmt.Scan(&name)

			input := command.Command{Name: "/Register", Args: []string{name}}
			command.Dispatch(input)
		case 0:
			return
		case -1:
			// 将来执行开发者模式，直接输入指令，以上帝视角操控玩家状态。
			input := command.Parser()
			command.Dispatch(input)
		default:
			continue
		}

		fmt.Println()
	}

}

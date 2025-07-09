package command

import (
	"fmt"
	"go-game-server/internal/match"
	"go-game-server/internal/player"
	"strconv"
)

// 接收解析好的命令，判断要调用哪个模块
func Dispatch(command Command) {
	switch command.Name {
	case "/Login":
		// 注意参数校验
		if len(command.Args) == 1 {
			// 获取ID
			id, err := strconv.Atoi(command.Args[0]) // 这里采用id登录
			if err == nil {
				player, ok := player.GetPlayerByID(id)
				if ok {
					player.Login()
				} else {
					// 此时没有通过id拿到玩家数据，表明该id未注册
					fmt.Println("请先注册")
				}
			} else {
				fmt.Printf("id转换数字失败。错误：%v\n", err)
			}
		} else {
			fmt.Printf("登录命令的参数不正确。 args = %v\n", command.Args)
		}
	case "/Logout":
		// 参数校验
		if len(command.Args) == 1 {
			id, err := strconv.Atoi(command.Args[0])
			if err == nil {
				player, ok := player.GetPlayerByID(id)
				if ok {
					player.Logout()
				} else {
					// 此时没有通过id拿到玩家数据，表明该id未注册
					fmt.Println("请先注册")
				}
			} else {
				fmt.Println("id转换数字失败")
			}
		} else {
			fmt.Println("登录命令的参数不正确")
		}
	case "/Register":
		// 参数校验
		if len(command.Args) == 1 {
			player.Register(command.Args[0])
		} else {
			fmt.Println("登录命令的参数不正确")
		}
	case "/status":
		// 参数校验
		if len(command.Args) == 1 {
			id, err := strconv.Atoi(command.Args[0])
			if err == nil {
				player.Status(id)
			} else {
				fmt.Println("id转换数字失败")
			}
		} else {
			fmt.Println("登录命令的参数不正确")
		}
	case "/match":
		// 参数校验
		if len(command.Args) == 1 {
			// 根据玩家id找到对应palyer对象
			id, err := strconv.Atoi(command.Args[0])
			if err == nil {
				player, err := player.QueryPlayer(id)
				if err == nil {
					match.Match(&player)
				} else {
					fmt.Print("该玩家不存在，请先注册！")
				}
			} else {
				fmt.Println("id转换数字失败")
			}
		} else {
			fmt.Println("登录命令的参数不正确")
		}
	default:
		fmt.Println("没有该命令！")
	} //case

}

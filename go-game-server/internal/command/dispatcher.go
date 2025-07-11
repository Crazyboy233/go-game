package command

import (
	"fmt"
	"go-game-server/internal/match"
	"go-game-server/internal/player"
	"strconv"
)

// 接收解析好的命令，判断要调用哪个模块
// return 是否执行成功
func Dispatch(command Command) bool {
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
					return true
				} else {
					// 此时没有通过id拿到玩家数据，表明该id未注册
					fmt.Println("请先注册！")
					return false
				}
			} else {
				fmt.Printf("id转换数字失败。错误：%v\n", err)
				return false
			}
		} else {
			fmt.Printf("登录命令的参数不正确。 args = %v\n", command.Args)
			return false
		}
	case "/Logout":
		// 参数校验
		if len(command.Args) == 1 {
			id, err := strconv.Atoi(command.Args[0])
			if err == nil {
				player, ok := player.GetPlayerByID(id)
				if ok {
					player.Logout()
					return true
				} else {
					// 此时没有通过id拿到玩家数据，表明该id未注册
					fmt.Println("该用户不存在，请先注册")
					return false
				}
			} else {
				fmt.Println("id转换数字失败")
				return false
			}
		} else {
			fmt.Println("登录命令的参数不正确")
			return false
		}
	case "/Register":
		// 参数校验
		if len(command.Args) == 1 {
			player.Register(command.Args[0])
			return true
		} else {
			fmt.Println("登录命令的参数不正确")
			return false
		}
	case "/status":
		// 参数校验
		if len(command.Args) == 1 {
			id, err := strconv.Atoi(command.Args[0])
			if err == nil {
				// TODO: 这里展示个人信息，要接收Status的返回值。
				player := player.Status(id)
				fmt.Println("===========================")
				fmt.Printf("账号: %d\n", player.GetId())
				fmt.Printf("名称: %v\n", player.GetName())
				fmt.Printf("在线状态: %v\n", player.IsOnline())
				fmt.Println("===========================")
			} else {
				fmt.Println("id转换数字失败")
				return false
			}
		} else {
			fmt.Println("登录命令的参数不正确")
			return false
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
					return false
				}
			} else {
				fmt.Println("id转换数字失败")
				return false
			}
		} else {
			fmt.Println("登录命令的参数不正确")
			return false
		}
	default:
		fmt.Println("没有该命令！")
		return false
	}

	return false
}

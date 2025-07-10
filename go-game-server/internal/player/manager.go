package player

import (
	"fmt"
)

var Manager = make(map[int]*Player)

func GetPlayerByID(id int) (*Player, bool) {
	p, ok := Manager[id]
	return p, ok
}

func QueryPlayer(id int) (Player, error) {
	if Manager == nil {
		return Player{}, fmt.Errorf("玩家管理器未初始化")
	}

	value, exists := Manager[id]
	if !exists {
		return Player{}, fmt.Errorf("查找的玩家不存在")
	}
	return *value, nil
}

// 注册玩家，id自动分配
// 注:玩家名称可以重复,只要注册,就认为是新号,而不是重复注册.
func Register(name string) *Player {
	// 方法1. map的大小+1就是下一个id值
	id := len(Manager) + 1
	player := &Player{id, name, false}
	Manager[id] = player
	fmt.Printf("注册成功，你的游戏账号为 %d\n", id)
	fmt.Println("请重新登陆！")
	return player

	// 方法2. 生成随机值，检查是否有该id。再赋值
}

// 查看玩家信息
func Status(id int) Player {
	player := Manager[id]
	return *player
}

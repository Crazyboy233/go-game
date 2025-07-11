package player

import "fmt"

// 玩家结构体与方法

type Player struct {
	id     int
	name   string
	online bool
}

// 登录
// return 是否登录成功
func (player *Player) Login() bool {
	// 先判断该玩家是否已经注册
	// 注：其实在分发的时候已经通过id查找过了，走到这里，用户一定存在
	_, ok := Manager[player.id]
	if !ok {
		// 该用户没注册
		return false
	}

	// 该用户已注册，将其在线状态改为登录状态
	player.online = true
	fmt.Println("登录成功")
	return true
}

// 登出
func (player *Player) Logout() {
	if player.online {
		player.online = false
	}
}

// 获取玩家id
func (player Player) GetId() int {
	return player.id
}

// 获取玩家名称
func (player Player) GetName() string {
	return player.name
}

// 判断玩家是否在线
func (player Player) IsOnline() bool {
	return player.online
}

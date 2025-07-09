package player

// 玩家结构体与方法

type Player struct {
	id     int
	name   string
	online bool
}

// 登录
func (player *Player) Login() {
	if !player.online {
		player.online = true
	}
}

// 登出
func (player *Player) Logout() {
	if player.online {
		player.online = false
	}
}

// 判断玩家是否在线
func (player Player) IsOnline() bool {
	return player.online
}

// 获取玩家名称
func (player Player) GetName() string {
	return player.name
}

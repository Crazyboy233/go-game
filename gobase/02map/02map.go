package main

import "fmt"

// 创建一个 map，用来保存玩家的在线状态（map[string]bool），然后模拟上线/下线/查询操作。
type Player struct {
	name   string
	online bool
}

func (p *Player) login() {
	p.online = true
}

func (p *Player) logout() {
	p.online = false
}

func (p *Player) query() string {
	if p.online {
		return "在线"
	}
	return "离线"
}

func main() {
	// 有点强行用map了，这里直接用slice存放*Player即可
	var players = make(map[string]*Player)
	players["Alice"] = &Player{name: "Alice"}
	players["Bob"] = &Player{name: "Bob"}

	players["Bob"].login()
	for name, player := range players {
		fmt.Printf("%s 状态:%s\n", name, player.query())
	}
}

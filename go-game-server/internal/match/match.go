package match

import (
	"fmt"
	"go-game-server/internal/player"
)

// 实现匹配机制

// 使用一个队列实现匹配(silce 实现)
var queue []*player.Player

func Match(p *player.Player) {
	// 判断是否已经在队列中，是否在线

	// 先判断传进来的玩家是否在线
	if !p.IsOnline() {
		// 如果不在线，直接提示用户先登录
		fmt.Println("该玩家不在线，请先登录")
		return
	}

	// 先判断队列是否为空，为空直接加入队列
	if len(queue) == 0 {
		queue = append(queue, p)
		return
	}

	// 走到这里说明玩家在线，且队列不为空。
	// 判断是否是玩家重复匹配
	for _, v := range queue {
		if v == p {
			// 已经在队列中了，提示用户并返回
			fmt.Println("已经在匹配啦，请稍等...")
			return
		} else {
			// 不在队列中，进行匹配，并将队列清空
			fmt.Printf("%s vs %s\n", v.GetName(), p.GetName())
			queue = make([]*player.Player, 0)
		}
	}
}

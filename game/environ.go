package game

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/golog"
)

var log = golog.New("server")

// Environ 主游戏环境
type Environ struct {
	Peer cellnet.GenericPeer
}

// NewEnviron 初始化
func NewEnviron() *Environ {
	return &Environ{}
}

// Update 更新游戏状态
func (env *Environ) Update() {
	// log.Debugln("Update")
}

// HandleEvent 处理客户端包
func (env *Environ) HandleEvent(ev cellnet.Event) {
	// log.Debugln("HandleEvent")
}

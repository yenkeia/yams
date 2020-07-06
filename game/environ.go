package game

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/golog"
	"github.com/yenkeia/yams/game/proto/client"
	"github.com/yenkeia/yams/game/proto/server"
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
	s := ev.Session()

	switch msg := ev.Message().(type) {
	case *cellnet.SessionAccepted: // 有新的连接
		s.Send(&server.Connected{})
	case *cellnet.SessionClosed: // 有连接断开
		// sessionClosed(s, msg)
	case *client.ClientVersion:
		s.Send(&server.ClientVersion{Result: 1})
	case *client.KeepAlive:
		s.Send(&server.KeepAlive{Time: 0})
	case *client.NewAccount:
		// newAccount(s, msg)
	case *client.ChangePassword:
		// changePassword(s, msg)
	case *client.Login:
		// login(s, msg)
	case *client.NewCharacter:
		// newCharacter(s, msg)
	case *client.DeleteCharacter:
		// deleteCharacter(s, msg)
	case *client.StartGame:
		// startGame(s, msg)
	case *client.LogOut:
		// logOut(s, msg)

	default:
		_ = msg

		// p, ok := g.GetPlayer(s, GAME)
		// if !ok {
		// 	return
		// }

		// handleEvent(p, g, ev, s)
	}
}

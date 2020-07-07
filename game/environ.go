package game

import (
	"fmt"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/golog"
	"github.com/jinzhu/gorm"
	"github.com/yenkeia/yams/game/proto/client"
	"github.com/yenkeia/yams/game/proto/server"
)

var (
	log  *golog.Logger
	db   *gorm.DB
	conf *config
)

func init() {
	conf = newConfig("../../configs/yams.yaml")
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.DB))
	if err != nil {
		panic(err)
	}
	log = golog.New("yams.game")
	_ = db
}

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
		newAccount(s, msg)
	case *client.ChangePassword:
		changePassword(s, msg)
	case *client.Login:
		login(s, msg)
	case *client.NewCharacter:
		newCharacter(s, msg)
	case *client.DeleteCharacter:
		deleteCharacter(s, msg)
	case *client.StartGame:
		startGame(s, msg)
	case *client.LogOut:
		logout(s, msg)
	default:
		_ = msg

		// 验证登陆状态

		// handleEvent(p, g, ev, s)
	}
}

func newAccount(s cellnet.Session, msg *client.NewAccount) {

}

func changePassword(s cellnet.Session, msg *client.ChangePassword) {

}

func login(s cellnet.Session, msg *client.Login) {

}

func newCharacter(s cellnet.Session, msg *client.NewCharacter) {

}

func deleteCharacter(s cellnet.Session, msg *client.DeleteCharacter) {

}

func startGame(s cellnet.Session, msg *client.StartGame) {

}

func logout(s cellnet.Session, msg *client.LogOut) {

}

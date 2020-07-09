package game

import (
	"fmt"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/golog"
	"github.com/jinzhu/gorm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/client"
	"github.com/yenkeia/yams/game/proto/server"
)

var log = golog.New("yams.game")
var sessionPlayer = make(map[int64]*player)
var db *gorm.DB

// Environ 主游戏环境
type Environ struct {
	Peer cellnet.GenericPeer
}

// NewEnviron 初始化
func NewEnviron(conf *Config) *Environ {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.DB))
	defer db.Close()
	if err != nil {
		panic(err)
	}
	return &Environ{}
}

// Update 更新游戏状态
func (env *Environ) Update() {
	// log.Debugln("Update")
}

// HandleEvent 处理客户端包
func (env *Environ) HandleEvent(e cellnet.Event) {
	s := e.Session()

	switch msg := e.Message().(type) {
	case *cellnet.SessionAccepted: // 有新的连接
		s.Send(&server.Connected{})
	case *cellnet.SessionClosed: // 有连接断开
		// sessionClosed(s, msg)
	case *client.ClientVersion:
		clientVersion(s, msg)
	case *client.KeepAlive:
		s.Send(&server.KeepAlive{Time: 0})
	case *client.NewAccount:
		newAccount(s, msg)
	case *client.ChangePassword:
		changePassword(s, msg)
	case *client.Login:
		login(s, msg, env)
	case *client.NewCharacter:
		newCharacter(s, msg, env)
	case *client.DeleteCharacter:
		deleteCharacter(s, msg)
	case *client.StartGame:
		startGame(s, msg)
	case *client.LogOut:
		logout(s, msg)
	default:
		if !checkGameStage(s, GAME) {
			return
		}
		p := sessionPlayer[s.ID()]
		handleEvent(p, e, s)
	}
}

func checkGameStage(s cellnet.Session, gameStage int) bool {
	player, ok := sessionPlayer[s.ID()]
	if !ok {
		return false
	}
	if player.gameStage != gameStage {
		return false
	}
	return true
}

func clientVersion(s cellnet.Session, msg *client.ClientVersion) {
	player := new(player)
	player.gameStage = LOGIN
	sessionPlayer[s.ID()] = player
	s.Send(&server.ClientVersion{Result: 1})
}

func newAccount(s cellnet.Session, msg *client.NewAccount) {
	if !checkGameStage(s, LOGIN) {
		return
	}
	res := uint8(0)
	a := new(orm.Account)
	db.Table("account").Where("username = ?", msg.UserName).Find(a)
	if a.ID == 0 {
		a.Username = msg.AccountID
		a.Password = msg.Password
		db.Table("account").Create(a)
		res = 8
	}
	s.Send(&server.NewAccount{Result: res})
}

func changePassword(s cellnet.Session, msg *client.ChangePassword) {
	if !checkGameStage(s, LOGIN) {
		return
	}
	res := uint8(5)
	a := new(orm.Account)
	db.Table("account").Where("username = ? AND password = ?", msg.AccountID, msg.CurrentPassword).Find(a)
	if a.ID != 0 {
		a.Password = msg.NewPassword
		db.Table("account").Model(a).Updates(orm.Account{Password: msg.NewPassword})
		res = 6
	}
	s.Send(&server.ChangePassword{Result: res})
}

func login(s cellnet.Session, msg *client.Login, env *Environ) {
	if !checkGameStage(s, LOGIN) {
		return
	}
	a := new(orm.Account)
	db.Table("account").Where("username = ? AND password = ?", msg.AccountID, msg.Password).Find(a)
	if a.ID == 0 {
		s.Send(&server.Login{Result: uint8(4)})
		return
	}
	res := new(server.LoginSuccess)
	res.Characters = nil // TODO 查询角色
	s.Send(res)
}

func newCharacter(s cellnet.Session, msg *client.NewCharacter, env *Environ) {
	if !checkGameStage(s, SELECT) {
		return
	}
	player := sessionPlayer[s.ID()]

	ac := make([]orm.AccountCharacter, 3)
	db.Table("account_character").Where("account_id = ?", player.accountID).Limit(3).Find(&ac)
	if len(ac) >= 3 {
		s.Send(&server.NewCharacter{Result: uint8(4)})
		return
	}

	// c := new(orm.Character)
	// TODO 判断角色名字是否重复

	res := new(server.NewCharacterSuccess)
	res.CharInfo = server.SelectInfo{
		Index:      1, // TODO 顺序 uint32(c.ID)
		Name:       msg.Name,
		Level:      1,          // uint16
		Class:      msg.Class,  // cm.MirClass
		Gender:     msg.Gender, // cm.MirGender
		LastAccess: 0,          // int6
	}
	s.Send(res)
}

func deleteCharacter(s cellnet.Session, msg *client.DeleteCharacter) {
	if !checkGameStage(s, SELECT) {
		return
	}
	c := new(orm.Character)
	db.Table("character").Where("id = ?", msg.CharacterIndex).Find(c)
	if c.ID == 0 {
		res := new(server.DeleteCharacter)
		res.Result = 4
		s.Send(res)
		return
	}
	db.Table("character").Delete(c)
	db.Table("account_character").Where("character_id = ?", c.ID).Delete(orm.Character{})
	res := new(server.DeleteCharacterSuccess)
	res.CharacterIndex = msg.CharacterIndex
	s.Send(res)
}

func startGame(s cellnet.Session, msg *client.StartGame) {
	if !checkGameStage(s, SELECT) {
		return
	}
	// TODO
}

func logout(s cellnet.Session, msg *client.LogOut) {

}

func handleEvent(p *player, e cellnet.Event, s cellnet.Session) {
	switch msg := e.Message().(type) {
	case *client.Turn:
		p.Turn(msg)
	case *client.Walk:
		p.Walk(msg)
	case *client.Run:
		p.Run(msg)
	case *client.Chat:
		p.Chat(msg)
	case *client.MoveItem:
		p.MoveItem(msg)
	case *client.StoreItem:
		p.StoreItem(msg)
	case *client.DepositRefineItem:
		p.DepositRefineItem(msg)
	case *client.RetrieveRefineItem:
		p.RetrieveRefineItem(msg)
	case *client.RefineCancel:
		p.RefineCancel(msg)
	case *client.RefineItem:
		p.RefineItem(msg)
	case *client.CheckRefine:
		p.CheckRefine(msg)
	case *client.ReplaceWedRing:
		p.ReplaceWedRing(msg)
	case *client.DepositTradeItem:
		p.DepositTradeItem(msg)
	case *client.RetrieveTradeItem:
		p.RetrieveTradeItem(msg)
	case *client.TakeBackItem:
		p.TakeBackItem(msg)
	case *client.MergeItem:
		p.MergeItem(msg)
	case *client.EquipItem:
		p.EquipItem(msg)
	case *client.RemoveItem:
		p.RemoveItem(msg)
	case *client.RemoveSlotItem:
		p.RemoveSlotItem(msg)
	case *client.SplitItem:
		p.SplitItem(msg)
	case *client.UseItem:
		p.UseItem(msg)
	case *client.DropItem:
		p.DropItem(msg)
	case *client.DropGold:
		p.DropGold(msg)
	case *client.PickUp:
		p.PickUp(msg)
	case *client.Inspect:
		p.Inspect(msg)
	// case *client.ChangeAMode:
	// 	p.ChangeAMode(msg)
	default:
		log.Debugln("default:", msg)
		//MessageQueue.Enqueue(string.Format("Invalid packet received. Index : {0}", p.Index));
	}
}

package game

import (
	"github.com/davyxu/cellnet"
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/client"
	"github.com/yenkeia/yams/game/proto/server"
)

type player struct {
	session         *cellnet.Session
	gameStage       int
	accountID       int
	objectID        int
	name            string
	nameColor       cm.Color
	currentMap      *mirMap
	currentLocation cm.Point
	bindLocation    cm.Point
	bindMap         *mirMap
	direction       cm.MirDirection
	hp              int
	mp              int
	maxHP           int
	maxMP           int
	level           int
	experience      int
	maxExperience   int
	guildName       string
	guildRankName   string
	class           cm.MirClass
	gender          cm.MirGender
	hair            int
	light           int
	gold            int
	inventory       *bag // 46
	equipment       *bag // 14
	questInventory  *bag // 40
	storage         *bag // 80
	trade           *bag // 10	交易框的索引是从上到下的，背包是从左到右
	attackMode      cm.AttackMode
	petMode         cm.PetMode
	allowGroup      bool
}

func (p *player) id() int {
	return p.objectID
}

func (p *player) enqueue(msg interface{}) {
	if msg == nil {
		log.Errorln("warning: enqueue nil message")
		return
	}
	(*p.session).Send(msg)
}

func (p *player) receiveChat(text string, typ cm.ChatType) {
	p.enqueue(&server.Chat{Message: text, Type: typ})
}

// FIXME
func (p *player) updateInfo(c *orm.Character) {
	p.gameStage = GAME
	p.objectID = 123 // TODO
	p.name = c.Name
	p.direction = cm.MirDirection(c.Direction)
	p.currentMap = env.maps[1] // TODO
	p.currentLocation = cm.NewPoint(int(c.CurrentLocationX), int(c.CurrentLocationY))
	p.bindLocation = cm.NewPoint(c.BindLocationX, c.BindLocationY)
	p.bindMap = env.maps[c.BindMapID]
	p.direction = cm.MirDirectionUp
	p.hp = c.HP
	p.mp = c.MP
	p.level = c.Level
	p.experience = c.Experience
	p.maxExperience = c.Experience + 100 // TODO
	p.guildName = ""                     // TODO
	p.guildRankName = ""                 // TODO
	p.class = cm.MirClass(c.Class)
	p.gender = cm.MirGender(c.Gender)
	p.hair = c.Hair
	p.light = 1 // TODO
	p.gold = c.Gold
	p.inventory = &bag{items: make([]*orm.UserItem, 46)}      // 46
	p.equipment = &bag{items: make([]*orm.UserItem, 14)}      // 14
	p.questInventory = &bag{items: make([]*orm.UserItem, 40)} // 40
	p.storage = &bag{items: make([]*orm.UserItem, 80)}        // 80
	p.trade = &bag{items: make([]*orm.UserItem, 10)}          // 10	交易框的索引是从上到下的，背包是从左到右
	p.attackMode = cm.AttackModeAll
	p.petMode = cm.PetModeBoth
	p.allowGroup = true
	/* TODO
	switch p.class {
	case cm.MirClassWarrior:
		p.maxHP = int(14.0 + (float32(p.level)/baseStats.HpGain+baseStats.HpGainRate+float32(p.level)/20.0)*float32(p.level))
		p.maxMP = int(11.0 + (float32(p.level) * 3.5) + (float32(p.level) * baseStats.MpGainRate))
	case cm.MirClassWizard:
		p.maxMP = int(13.0 + (float32(p.level/5.0+2.0) * 2.2 * float32(p.level)) + (float32(p.level) * baseStats.MpGainRate))
	case cm.MirClassTaoist:
		p.maxMP = int((13 + float32(p.level)/8.0*2.2*float32(p.level)) + (float32(p.level) * baseStats.MpGainRate))
	}
	*/
}

func (p *player) turn(msg *client.Turn) {
	p.direction = msg.Direction
	p.enqueue(&server.UserLocation{Location: p.currentLocation, Direction: p.direction})
}

func (p *player) walk(msg *client.Walk) {
	p.direction = msg.Direction
	p.currentLocation = p.currentLocation.NextPoint(msg.Direction, 1)
	p.enqueue(&server.UserLocation{Location: p.currentLocation, Direction: p.direction})
}

func (p *player) run(msg *client.Run) {
	p.direction = msg.Direction
	p.currentLocation = p.currentLocation.NextPoint(msg.Direction, 2)
	p.enqueue(&server.UserLocation{Location: p.currentLocation, Direction: p.direction})
}

func (p *player) chat(msg *client.Chat)                             {}
func (p *player) moveItem(msg *client.MoveItem)                     {}
func (p *player) storeItem(msg *client.StoreItem)                   {}
func (p *player) depositRefineItem(msg *client.DepositRefineItem)   {}
func (p *player) retrieveRefineItem(msg *client.RetrieveRefineItem) {}
func (p *player) refineCancel(msg *client.RefineCancel)             {}
func (p *player) refineItem(msg *client.RefineItem)                 {}
func (p *player) checkRefine(msg *client.CheckRefine)               {}
func (p *player) replaceWedRing(msg *client.ReplaceWedRing)         {}
func (p *player) depositTradeItem(msg *client.DepositTradeItem)     {}
func (p *player) retrieveTradeItem(msg *client.RetrieveTradeItem)   {}
func (p *player) takeBackItem(msg *client.TakeBackItem)             {}
func (p *player) mergeItem(msg *client.MergeItem)                   {}
func (p *player) equipItem(msg *client.EquipItem)                   {}
func (p *player) removeItem(msg *client.RemoveItem)                 {}
func (p *player) removeSlotItem(msg *client.RemoveSlotItem)         {}
func (p *player) splitItem(msg *client.SplitItem)                   {}
func (p *player) useItem(msg *client.UseItem)                       {}
func (p *player) dropItem(msg *client.DropItem)                     {}
func (p *player) dropGold(msg *client.DropGold)                     {}
func (p *player) pickUp(msg *client.PickUp)                         {}
func (p *player) inspect(msg *client.Inspect)                       {}

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
	accountID       int // account.ID
	characterID     int // character.ID 保存数据库用
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

func (p *player) getObjectID() int {
	return p.objectID
}

func (p *player) getPosition() cm.Point {
	return p.currentLocation
}

func (p *player) enqueue(msg interface{}) {
	if msg == nil {
		log.Errorln("warning: enqueue nil message")
		return
	}
	(*p.session).Send(msg)
}

// TODO
func (p *player) enqueueItemInfos() {

}

// TODO
func (p *player) enqueueQuestInfo() {

}

func (p *player) enqueueAreaObjects(g1, g2 *aoiGrid) {
	area1 := make([]*aoiGrid, 0)
	if g1 != nil {
		area1 = p.currentMap.aoi.getSurroundGridsByGid(g1.gID)
	}
	area2 := p.currentMap.aoi.getSurroundGridsByGid(g2.gID)
	send := make(map[int]bool)
	for x := range area2 {
		send[area2[x].gID] = true
		for y := range area1 {
			if area1[y].gID == area2[x].gID {
				send[area2[x].gID] = false
			}
		}
	}
	for x := range area2 {
		if send[area2[x].gID] {
			objs := env.getMapObjects(area2[x].getObjectIDs())
			for _, obj := range objs {
				if obj.getObjectID() == p.objectID {
					continue
				}
				p.enqueueMapObject(obj)
			}
		}
	}
	drop := make(map[int]bool)
	for x := range area1 {
		drop[area1[x].gID] = true
		for y := range area2 {
			if area1[x].gID == area2[y].gID {
				drop[area2[y].gID] = false
			}
		}
	}
	for x := range area1 {
		if drop[area1[x].gID] {
			objs := env.getMapObjects(area1[x].getObjectIDs())
			for _, obj := range objs {
				if obj.getObjectID() == p.objectID {
					continue
				}
				p.enqueue(&server.ObjectRemove{ObjectID: uint32(obj.getObjectID())})
			}
		}
	}
}

// TODO
func (p *player) getObjectPlayer() *server.ObjectPlayer {
	return &server.ObjectPlayer{
		ObjectID:         uint32(p.objectID),      // uint32
		Name:             p.name,                  // string
		GuildName:        "",                      // string
		GuildRankName:    "",                      // string
		NameColor:        cm.ColorWhite.ToInt32(), // int32 // = Color.FromArgb(reader.ReadInt32());
		Class:            p.class,                 // cm.MirClass
		Gender:           p.gender,                // cm.MirGender
		Level:            uint16(p.level),         // uint16
		Location:         p.currentLocation,       // cm.Point
		Direction:        p.direction,             // cm.MirDirection
		Hair:             uint8(p.hair),           // uint8
		Light:            uint8(p.light),          // uint8
		Weapon:           0,                       // int16
		WeaponEffect:     0,                       // int16
		Armour:           0,                       // int16
		Poison:           0,                       // cm.PoisonType // = (PoisonType)reader.ReadUInt16()
		Dead:             false,                   // bool
		Hidden:           false,                   // bool
		Effect:           0,                       // cm.SpellEffect // = (SpellEffect)reader.ReadByte()
		WingEffect:       0,                       // uint8
		Extra:            false,                   // bool
		MountType:        0,                       // int16
		RidingMount:      false,                   // bool
		Fishing:          false,                   // bool
		TransformType:    0,                       // int16
		ElementOrbEffect: 0,                       // uint32
		ElementOrbLvl:    0,                       // uint32
		ElementOrbMax:    0,                       // uint32
		Buffs:            make([]cm.BuffType, 0),  // []cm.BuffType
		LevelEffects:     0,                       // cm.LevelEffects
	}
}

// TODO
func (p *player) enqueueMapObject(obj mapObject) {
	switch o := obj.(type) {
	case *player:
		p.enqueue(o.getObjectPlayer())
	case *npc:
		p.enqueue(&server.ObjectNPC{
			ObjectID:  uint32(o.objectID),
			Name:      o.name,
			NameColor: cm.ColorWhite.ToInt32(),
			Image:     uint16(o.info.Image),
			Color:     0,
			Location:  o.getPosition(),
			Direction: o.direction,
			QuestIDs:  make([]int32, 0),
		})
	case *monster:
		p.enqueue(&server.ObjectMonster{
			ObjectID:          uint32(o.objectID),
			Name:              o.info.Name,
			NameColor:         cm.ColorWhite.ToInt32(),
			Location:          o.location,
			Image:             cm.Monster(o.info.Image),
			Direction:         o.direction,
			Effect:            uint8(o.info.Effect),
			AI:                uint8(o.info.AI),
			Light:             uint8(o.info.Light),
			Dead:              o.isDead,
			Skeleton:          o.isSkeleton,
			Poison:            o.poison,
			Hidden:            o.isHidden,
			ShockTime:         0,     // TODO
			BindingShotCenter: false, // TODO
			Extra:             false, // TODO
			ExtraByte:         0,     // TODO
		})
	}

}

func (p *player) broadcast(msg interface{}) {
	mp := env.maps[p.currentMap.info.ID]
	mp.broadcast(p.currentLocation, msg)
}

func (p *player) receiveChat(text string, typ cm.ChatType) {
	p.enqueue(&server.Chat{Message: text, Type: typ})
}

// FIXME
func (p *player) updateInfo(c *orm.Character) {
	p.gameStage = GAME
	p.objectID = env.newObjectID()
	p.characterID = c.ID
	p.name = c.Name
	p.direction = cm.MirDirection(c.Direction)
	p.currentMap = env.maps[c.CurrentMapID]
	p.currentLocation = cm.NewPoint(int(c.CurrentLocationX), int(c.CurrentLocationY))
	p.bindLocation = cm.NewPoint(c.BindLocationX, c.BindLocationY)
	p.bindMap = env.maps[c.BindMapID]
	p.direction = cm.MirDirection(c.Direction)
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

// TODO
func (p *player) refreshStats() {

}

func (p *player) turn(msg *client.Turn) {
	p.direction = msg.Direction
	p.enqueue(&server.UserLocation{Location: p.currentLocation, Direction: p.direction})
}

func (p *player) walk(msg *client.Walk) {
	p.direction = msg.Direction
	p.currentMap.updateObject(p, p.currentLocation.NextPoint(msg.Direction, 1))
	p.currentLocation = p.currentLocation.NextPoint(msg.Direction, 1)
	p.enqueue(&server.UserLocation{Location: p.currentLocation, Direction: p.direction})
}

func (p *player) run(msg *client.Run) {
	p.direction = msg.Direction
	p.currentMap.updateObject(p, p.currentLocation.NextPoint(msg.Direction, 2))
	p.currentLocation = p.currentLocation.NextPoint(msg.Direction, 2)
	p.enqueue(&server.UserLocation{Location: p.currentLocation, Direction: p.direction})
}

func (p *player) chat(msg *client.Chat) {
	res := &server.ObjectChat{
		ObjectID: uint32(p.objectID),
		Text:     p.name + ":" + msg.Message,
		Type:     cm.ChatTypeNormal,
	}
	p.enqueue(res)
	p.broadcast(res)
}

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

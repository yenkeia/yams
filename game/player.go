package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/davyxu/cellnet"
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/client"
	"github.com/yenkeia/yams/game/proto/server"
)

type player struct {
	baseObject
	session           *cellnet.Session
	actionList        *actionList
	gameStage         int
	accountID         int // account.ID
	characterID       int // character.ID 保存数据库用
	bindLocation      cm.Point
	bindMapID         int
	hp                int
	mp                int
	maxHP             int
	maxMP             int
	level             int
	experience        int
	maxExperience     int
	guildName         string
	guildRankName     string
	class             cm.MirClass
	gender            cm.MirGender
	hair              int
	light             int
	looksWeapon       int
	looksWeaponEffect int
	looksArmour       int
	looksWings        int
	gold              int
	inventory         *bag // 46
	equipment         *bag // 14
	questInventory    *bag // 40
	storage           *bag // 80
	trade             *bag // 10	交易框的索引是从上到下的，背包是从左到右
	attackMode        cm.AttackMode
	petMode           cm.PetMode
	allowGroup        bool
	sendedItemInfoIDs []int
	dead              bool
	callingNPC        int // obejctID
	callingNPCKey     string
	minAC             int // 物理防御力
	maxAC             int
	minMAC            int // 魔法防御力
	maxMAC            int
	minDC             int // 攻击力
	maxDC             int
	minMC             int // 魔法力
	maxMC             int
	minSC             int // 道术力
	maxSC             int
}

func (p *player) String() string {
	return fmt.Sprintf("玩家: %s, 等级: %d, 位置: %s", p.name, p.level, p.location)
}

func (p *player) getObjectID() int {
	return p.objectID
}

func (p *player) getPosition() cm.Point {
	return p.location
}

func (p *player) update(now time.Time) {
	p.actionList.execute()
}

func (p *player) getAttackPower(min, max int) int {
	if min < 0 {
		min = 0
	}
	if max < min {
		max = min
	}
	// TODO luck
	return cm.RandomInt(min, max)
}

func (p *player) enqueue(msg interface{}) {
	if msg == nil {
		log.Errorln("warning: enqueue nil message")
		return
	}
	(*p.session).Send(msg)
}

func (p *player) enqueueItemInfo(itemID int) {
	for _, id := range p.sendedItemInfoIDs {
		if id == itemID {
			return
		}
	}
	item := gdb.itemInfoMap[itemID]
	if item == nil {
		return
	}
	p.enqueue(&server.NewItemInfo{Info: item.ToServerItemInfo()})
	p.sendedItemInfoIDs = append(p.sendedItemInfoIDs, itemID)
}

func (p *player) enqueueItemInfos() {
	itemInfos := make([]*orm.ItemInfo, 0)
	for _, v := range p.inventory.items {
		if v != nil {
			itemInfos = append(itemInfos, gdb.itemInfoMap[v.info.ID])
		}
	}
	for _, v := range p.equipment.items {
		if v != nil {
			itemInfos = append(itemInfos, gdb.itemInfoMap[v.info.ID])
		}
	}
	for _, v := range p.questInventory.items {
		if v != nil {
			itemInfos = append(itemInfos, gdb.itemInfoMap[v.info.ID])
		}
	}
	for i := range itemInfos {
		p.enqueueItemInfo(itemInfos[i].ID)
	}
}

// TODO
func (p *player) enqueueQuestInfo() {

}

func (p *player) enqueueAreaObjects(g1, g2 *aoiGrid) {
	area1 := make([]*aoiGrid, 0)
	mp := env.maps[p.mapID]
	if g1 != nil {
		area1 = mp.aoi.getSurroundGridsByGid(g1.gID)
	}
	area2 := mp.aoi.getSurroundGridsByGid(g2.gID)
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
		Location:         p.location,              // cm.Point
		Direction:        p.direction,             // cm.MirDirection
		Hair:             uint8(p.hair),           // uint8
		Light:            uint8(p.light),          // uint8
		Weapon:           int16(p.looksWeapon),
		WeaponEffect:     int16(p.looksWeaponEffect),
		Armour:           int16(p.looksArmour),
		Poison:           0,                      // cm.PoisonType // = (PoisonType)reader.ReadUInt16()
		Dead:             false,                  // bool
		Hidden:           false,                  // bool
		Effect:           0,                      // cm.SpellEffect // = (SpellEffect)reader.ReadByte()
		WingEffect:       0,                      // uint8
		Extra:            false,                  // bool
		MountType:        0,                      // int16
		RidingMount:      false,                  // bool
		Fishing:          false,                  // bool
		TransformType:    0,                      // int16
		ElementOrbEffect: 0,                      // uint32
		ElementOrbLvl:    0,                      // uint32
		ElementOrbMax:    0,                      // uint32
		Buffs:            make([]cm.BuffType, 0), // []cm.BuffType
		LevelEffects:     0,                      // cm.LevelEffects
	}
}

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
	case *item:
		p.enqueue(o.getItemObjectInfo())
	}
}

func (p *player) broadcast(msg interface{}) {
	env.maps[p.mapID].broadcast(p.location, msg, p.objectID)
}

func (p *player) broadcastPlayerUpdate() {
	p.broadcast(&server.PlayerUpdate{
		ObjectID:     uint32(p.objectID),
		Light:        uint8(p.light),
		Weapon:       int16(p.looksWeapon),
		WeaponEffect: int16(p.looksWeaponEffect),
		Armour:       int16(p.looksArmour),
		WingEffect:   uint8(p.looksWings),
	})
}

func (p *player) receiveChat(text string, typ cm.ChatType) {
	p.enqueue(&server.Chat{Message: text, Type: typ})
}

// FIXME
func (p *player) updateInfo(c *orm.Character) {
	p.actionList = newActionList()
	p.gameStage = GAME
	p.objectID = env.newObjectID()
	p.characterID = c.ID
	p.name = c.Name
	p.direction = cm.MirDirection(c.Direction)
	p.mapID = c.CurrentMapID
	p.location = cm.NewPoint(int(c.CurrentLocationX), int(c.CurrentLocationY))
	p.bindLocation = cm.NewPoint(c.BindLocationX, c.BindLocationY)
	p.bindMapID = c.BindMapID
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
	p.inventory = bagLoadFromDB(c.ID, cm.UserItemTypeInventory, 46)           // 46
	p.equipment = bagLoadFromDB(c.ID, cm.UserItemTypeEquipment, 14)           // 14
	p.questInventory = bagLoadFromDB(c.ID, cm.UserItemTypeQuestInventory, 40) // 40
	p.storage = bagLoadFromDB(c.ID, cm.UserItemTypeStorage, 80)               // 80
	p.trade = bagLoadFromDB(c.ID, cm.UserItemTypeTrade, 10)                   // 10	交易框的索引是从上到下的，背包是从左到右
	p.attackMode = cm.AttackModeAll
	p.petMode = cm.PetModeBoth
	p.allowGroup = true
	p.sendedItemInfoIDs = make([]int, 0)
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

func (p *player) updateConcentration() {
	p.enqueue(&server.SetConcentration{
		ObjectID:    uint32(p.accountID),
		Enabled:     false,
		Interrupted: false,
	})
	p.broadcast(&server.SetObjectConcentration{
		ObjectID:    uint32(p.accountID),
		Enabled:     false,
		Interrupted: false,
	})
}

// TODO
func (p *player) refreshStats() {

}

// TODO
func (p *player) refreshBagWeight() {

}

func (p *player) turn(msg *client.Turn) {
	p.direction = msg.Direction
	p.enqueue(&server.UserLocation{Location: p.location, Direction: p.direction})
}

func (p *player) walk(msg *client.Walk) {
	mp := env.maps[p.mapID]
	p.direction = msg.Direction
	mp.updateObject(p, p.location.NextPoint(msg.Direction, 1))
	p.location = p.location.NextPoint(msg.Direction, 1)
	p.enqueue(&server.UserLocation{Location: p.location, Direction: p.direction})
}

func (p *player) run(msg *client.Run) {
	mp := env.maps[p.mapID]
	p.direction = msg.Direction
	mp.updateObject(p, p.location.NextPoint(msg.Direction, 2))
	p.location = p.location.NextPoint(msg.Direction, 2)
	p.enqueue(&server.UserLocation{Location: p.location, Direction: p.direction})
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

func (p *player) getUserItemByObjectID(mirGridType cm.MirGridType, objectID int) (index int, item *userItem) {
	var arr []*userItem
	switch mirGridType {
	case cm.MirGridTypeInventory:
		arr = p.inventory.items
	case cm.MirGridTypeEquipment:
		arr = p.equipment.items
	case cm.MirGridTypeStorage:
		arr = p.storage.items
	default:
		panic("error mirGridType")
	}
	for i, v := range arr {
		if v != nil && v.objectID == objectID {
			return i, v
		}
	}
	return -1, nil
}

func (p *player) moveItem(msg *client.MoveItem) {
	res := &server.MoveItem{
		Grid:    msg.Grid,
		From:    msg.From,
		To:      msg.To,
		Success: false,
	}
	var err error
	switch msg.Grid {
	case cm.MirGridTypeInventory:
		err = p.inventory.move(int(msg.From), int(msg.To))
	case cm.MirGridTypeStorage:
		err = p.storage.move(int(msg.From), int(msg.To))
	case cm.MirGridTypeTrade:
		err = p.trade.move(int(msg.From), int(msg.To))
		p.tradeItem()
	case cm.MirGridTypeRefine:
		// TODO
	}
	if err != nil {
		p.receiveChat(err.Error(), cm.ChatTypeSystem)
	} else {
		res.Success = true
	}
	p.enqueue(res)
}

func (p *player) storeItem(msg *client.StoreItem)                   {}
func (p *player) depositRefineItem(msg *client.DepositRefineItem)   {}
func (p *player) retrieveRefineItem(msg *client.RetrieveRefineItem) {}

func (p *player) refineCancel(msg *client.RefineCancel) {
	p.callingNPC = 0
}

func (p *player) refineItem(msg *client.RefineItem)               {}
func (p *player) checkRefine(msg *client.CheckRefine)             {}
func (p *player) replaceWedRing(msg *client.ReplaceWedRing)       {}
func (p *player) depositTradeItem(msg *client.DepositTradeItem)   {}
func (p *player) retrieveTradeItem(msg *client.RetrieveTradeItem) {}
func (p *player) takeBackItem(msg *client.TakeBackItem)           {}
func (p *player) mergeItem(msg *client.MergeItem)                 {}

// 穿装备
func (p *player) equipItem(msg *client.EquipItem) {
	mirGridType := msg.Grid
	to := msg.To
	id := msg.UniqueID
	res := &server.EquipItem{
		Grid:     mirGridType,
		UniqueID: id,
		To:       to,
		Success:  false,
	}
	index, item := p.getUserItemByObjectID(mirGridType, int(id))
	if item == nil {
		p.enqueue(res)
		return
	}
	var err error
	switch mirGridType {
	case cm.MirGridTypeInventory:
		err = p.inventory.moveTo(index, int(to), p.equipment)
	case cm.MirGridTypeStorage:
		err = p.inventory.moveTo(index, int(to), p.storage)
	}
	if err != nil {
		p.enqueue(res)
		return
	}
	res.Success = true
	p.refreshStats()
	p.enqueue(res)
	p.updateConcentration()
	p.broadcastPlayerUpdate()
}

// 卸下装备
func (p *player) removeItem(msg *client.RemoveItem) {
	mirGridType := msg.Grid
	id := msg.UniqueID
	to := msg.To
	res := &server.RemoveItem{
		Grid:     mirGridType,
		UniqueID: id,
		To:       to,
		Success:  false,
	}
	index, item := p.getUserItemByObjectID(cm.MirGridTypeEquipment, int(id))
	if item == nil {
		p.enqueue(res)
		return
	}
	switch mirGridType {
	case cm.MirGridTypeInventory:
		if err := p.equipment.moveTo(index, int(msg.To), p.inventory); err != nil {
			p.receiveChat(err.Error(), cm.ChatTypeSystem)
			p.enqueue(res)
			return
		}
	case cm.MirGridTypeStorage:
		// TODO
		// if !util.StringEqualFold(p.CallingNPCPage, StorageKey) {
		// 	p.Enqueue(res)
		// 	return
		// }
		// p.Equipment.MoveTo(index, int(msg.To), p.Storage)
		p.receiveChat("没实现这个功能", cm.ChatTypeSystem)
		p.enqueue(res)
		return
	}
	res.Success = true
	p.refreshStats()
	p.enqueue(res)
	p.updateConcentration()
	p.broadcastPlayerUpdate()
}

func (p *player) removeSlotItem(msg *client.RemoveSlotItem) {}
func (p *player) splitItem(msg *client.SplitItem)           {}
func (p *player) useItem(msg *client.UseItem)               {}

func (p *player) dropItem(msg *client.DropItem) {
	res := &server.DropItem{
		UniqueID: msg.UniqueID,
		Count:    msg.Count,
		Success:  false,
	}
	count := int(msg.Count)
	index, userItem := p.getUserItemByObjectID(cm.MirGridTypeInventory, int(msg.UniqueID))
	if userItem == nil || count > userItem.count {
		p.enqueue(res)
		return
	}
	var err error
	obj := newItem(p.mapID, p.location, userItem)
	if err = obj.drop(p.location, 1); err != nil {
		p.receiveChat(err.Error(), cm.ChatTypeSystem)
		p.enqueue(res)
		return
	}
	if int(msg.Count) >= userItem.count {
		err = p.inventory.set(index, nil)
	} else {
		p.inventory.setCount(index, userItem.count-count)
	}
	res.Success = true
	if err != nil {
		res.Success = false
		log.Errorln(err.Error())
	}
	p.refreshBagWeight()
	p.enqueue(res)
}

func (p *player) dropGold(msg *client.DropGold) {
	amount := int(msg.Amount)
	if p.gold < amount {
		return
	}
	obj := newItemGold(p.mapID, p.location, amount)
	if err := obj.drop(p.location, 3); err != nil {
		p.receiveChat(err.Error(), cm.ChatTypeSystem)
		return
	}
	p.takeGold(amount)
}

// takeGold 玩家扣钱
func (p *player) takeGold(amount int) {
	if amount > p.gold {
		p.gold = 0
	} else {
		p.gold -= amount
	}
	pdb.syncGold(p)
	p.enqueue(&server.LoseGold{Gold: uint32(amount)})
}

// gainGold 玩家加钱
func (p *player) gainGold(amount int) {
	p.gold += amount
	pdb.syncGold(p)
	p.enqueue(&server.GainedGold{Gold: uint32(amount)})
}

func (p *player) gainItem(ui *userItem) (res bool) {
	// ui.soulBoundID = p.characterID
	ui.soulBoundID = -1

	if ui.info.StackSize > 1 {
		for i, v := range p.inventory.items {
			if v == nil || ui.info != v.info || v.count > ui.info.StackSize {
				continue
			}
			if ui.count+v.count <= ui.info.StackSize {
				p.inventory.setCount(i, v.count+ui.count)
				p.enqueue(&server.GainedItem{Item: ui.serverUserItem()})
				return true
			}
			p.inventory.setCount(i, v.count+ui.count)
			ui.count -= ui.info.StackSize - v.count
		}
	}

	i, j := 0, 46
	if ui.info.Type == cm.ItemTypePotion ||
		ui.info.Type == cm.ItemTypeScroll ||
		(ui.info.Type == cm.ItemTypeScript && ui.info.Effect == 1) {
		i = 0
		j = 4
	} else if ui.info.Type == cm.ItemTypeAmulet {
		i = 4
		j = 6
	} else {
		i = 6
		j = 46
	}
	for i < j {
		if p.inventory.items[i] != nil {
			i++
			continue
		}
		p.inventory.set(i, ui)
		// p.Inventory.Items[i] = ui
		p.enqueueItemInfo(ui.info.ID)
		p.enqueue(&server.GainedItem{Item: ui.serverUserItem()})
		p.refreshBagWeight()
		return true
	}
	i = 0
	for i < 46 {
		if p.inventory.items[i] != nil {
			i++
			continue
		}
		p.inventory.set(i, ui)
		p.enqueueItemInfo(ui.info.ID)
		p.enqueue(&server.GainedItem{Item: ui.serverUserItem()})
		p.refreshBagWeight()
		return true
	}
	p.receiveChat("没有合适的格子放置物品", cm.ChatTypeSystem)
	return false
}

func (p *player) pickUp(msg *client.PickUp) {
	if p.dead {
		return
	}
	mp := env.maps[p.mapID]
	c := mp.getCell(p.location)
	if c == nil {
		return
	}
	items := make([]*item, 0)
	for it := c.objects.Front(); it != nil; it = it.Next() {
		if i, ok := it.Value.(*item); ok {
			if i.ui == nil {
				p.gainGold(i.gold)
				items = append(items, i)
			} else {
				if p.gainItem(i.ui) {
					items = append(items, i)
				}
			}
		}
	}
	for i := range items {
		item := items[i]
		mp.deleteObject(item)
		delete(env.items, item.objectID)
		item.broadcast(&server.ObjectRemove{ObjectID: uint32(item.objectID)})
	}
}

func (p *player) inspect(msg *client.Inspect)         {}
func (p *player) changeAMode(msg *client.ChangeAMode) {}
func (p *player) changePMode(msg *client.ChangePMode) {}
func (p *player) changeTrade(msg *client.ChangeTrade) {}

func (p *player) attack(msg ...interface{}) {
	p.direction = msg[0].(*client.Attack).Direction
	p.enqueue(&server.UserLocation{Location: p.location, Direction: p.direction})
	p.broadcast(&server.ObjectAttack{
		ObjectID:  uint32(p.objectID),  // uint32
		LocationX: int32(p.location.X), // int32
		LocationY: int32(p.location.Y), // int32
		Direction: p.direction,         // cm.MirDirection
		Spell:     cm.SpellNone,        // cm.Spell
		Level:     0,                   // uint8
		Type:      0,                   // uint8
	})
	damageBase := p.getAttackPower(p.minDC, p.maxDC)
	damageFinal := damageBase          // TODO
	defence := cm.DefenceTypeACAgility // TODO
	mp := env.maps[p.mapID]
	cell := mp.getCell(p.location.NextPoint(p.direction, 1))
	if cell.attribute != cm.CellAttributeWalk {
		return
	}
	for e := cell.objects.Front(); e != nil; e = e.Next() {
		obj := e.Value.(attackableObject)
		p.actionList.pushDelayAction(cm.DelayedTypeDamage, 300, func() { p.completeAttack(obj, damageFinal, defence, true) })
	}
}

func (p *player) completeAttack(args ...interface{}) {
	target := args[0].(attackableObject)
	damage := args[1].(int)
	defence := args[2].(cm.DefenceType)
	damageWeapon := args[3].(bool)
	if target == nil || !target.isAttackTarget(p) { // || target.CurrentMap != CurrentMap || target.Node == nil) {
		return
	}
	if target.attacked(p, damage, defence, damageWeapon) <= 0 {
		return
	}
	/* TODO
	//Level Fencing / SpiritSword
	for _, magic := range p.Magics {
		switch magic.Spell {
		case cm.SpellFencing, cm.SpellSpiritSword:
			p.LevelMagic(magic)
			break
		}
	}
	*/
}

func (p *player) rangeAttack(msg *client.RangeAttack) {}
func (p *player) harvest(msg *client.Harvest)         {}

func (p *player) callNPC(msg *client.CallNPC) {
	// fmt.Println("->", msg.Key) // [@Main]
	// TODO
	// 判断玩家位置
	n, ok := env.npcs[int(msg.ObjectID)]
	if !ok {
		return
	}
	key := strings.ToUpper(msg.Key)
	say, err := n.script.call(key, n, p)
	if err != nil {
		log.Warnf("NPC 脚本执行失败: %d %s %s\n", n.objectID, key, err.Error())
	}
	log.Debugf("callNPC: %s %d, key: %s", n.name, n.objectID, key)
	p.callingNPC = int(msg.ObjectID)
	p.callingNPCKey = key
	p.enqueue(&server.NPCResponse{Page: replaceTemplates(n, p, say)})
	n.processSpecial(p, key)
}

// TODO
func replaceTemplates(n *npc, p *player, say []string) []string {
	say = say[1:]
	return say
}

func (p *player) talkMonsterNPC(msg *client.TalkMonsterNPC) {}

// 从 NPC 买东西
func (p *player) buyItem(msg *client.BuyItem) {
	n, ok := env.npcs[p.callingNPC]
	if !ok {
		return
	}
	var item *userItem
	for _, good := range n.goods {
		if good.objectID == int(msg.ItemIndex) {
			item = good
			break
		}
	}
	count := int(msg.Count)
	if item == nil || msg.Count == 0 || count > item.info.StackSize {
		return
	}
	price := item.price()
	if price > p.gold {
		p.receiveChat("金币不足", cm.ChatTypeSystem)
		return
	}
	newUserItem := newUserItem(item.info)
	if p.gainItem(newUserItem) {
		p.takeGold(price)
	}
}

func (p *player) craftItem(msg *client.CraftItem)                                 {}
func (p *player) sellItem(msg *client.SellItem)                                   {}
func (p *player) repairItem(msg *client.RepairItem)                               {}
func (p *player) buyItemBack(msg *client.BuyItemBack)                             {}
func (p *player) sRepairItem(msg *client.SRepairItem)                             {}
func (p *player) magicKey(msg *client.MagicKey)                                   {}
func (p *player) magic(msg *client.Magic)                                         {}
func (p *player) switchGroup(msg *client.SwitchGroup)                             {}
func (p *player) addMember(msg *client.AddMember)                                 {}
func (p *player) delMember(msg *client.DelMember)                                 {}
func (p *player) groupInvite(msg *client.GroupInvite)                             {}
func (p *player) townRevive(msg *client.TownRevive)                               {}
func (p *player) spellToggle(msg *client.SpellToggle)                             {}
func (p *player) consignItem(msg *client.ConsignItem)                             {}
func (p *player) marketSearch(msg *client.MarketSearch)                           {}
func (p *player) marketRefresh(msg *client.MarketRefresh)                         {}
func (p *player) marketPage(msg *client.MarketPage)                               {}
func (p *player) marketBuy(msg *client.MarketBuy)                                 {}
func (p *player) marketGetBack(msg *client.MarketGetBack)                         {}
func (p *player) requestUserName(msg *client.RequestUserName)                     {}
func (p *player) requestChatItem(msg *client.RequestChatItem)                     {}
func (p *player) editGuildMember(msg *client.EditGuildMember)                     {}
func (p *player) editGuildNotice(msg *client.EditGuildNotice)                     {}
func (p *player) guildInvite(msg *client.GuildInvite)                             {}
func (p *player) requestGuildInfo(msg *client.RequestGuildInfo)                   {}
func (p *player) guildNameReturn(msg *client.GuildNameReturn)                     {}
func (p *player) guildStorageGoldChange(msg *client.GuildStorageGoldChange)       {}
func (p *player) guildStorageItemChange(msg *client.GuildStorageItemChange)       {}
func (p *player) guildWarReturn(msg *client.GuildWarReturn)                       {}
func (p *player) marriageRequest(msg *client.MarriageRequest)                     {}
func (p *player) marriageReply(msg *client.MarriageReply)                         {}
func (p *player) changeMarriage(msg *client.ChangeMarriage)                       {}
func (p *player) divorceRequest(msg *client.DivorceRequest)                       {}
func (p *player) divorceReply(msg *client.DivorceReply)                           {}
func (p *player) addMentor(msg *client.AddMentor)                                 {}
func (p *player) mentorReply(msg *client.MentorReply)                             {}
func (p *player) allowMentor(msg *client.AllowMentor)                             {}
func (p *player) cancelMentor(msg *client.CancelMentor)                           {}
func (p *player) tradeRequest(msg *client.TradeRequest)                           {}
func (p *player) tradeGold(msg *client.TradeGold)                                 {}
func (p *player) tradeReply(msg *client.TradeReply)                               {}
func (p *player) tradeConfirm(msg *client.TradeConfirm)                           {}
func (p *player) tradeCancel(msg *client.TradeCancel)                             {}
func (p *player) tradeItem()                                                      {}
func (p *player) equipSlotItem(msg *client.EquipSlotItem)                         {}
func (p *player) fishingCast(msg *client.FishingCast)                             {}
func (p *player) fishingChangeAutocast(msg *client.FishingChangeAutocast)         {}
func (p *player) acceptQuest(msg *client.AcceptQuest)                             {}
func (p *player) finishQuest(msg *client.FinishQuest)                             {}
func (p *player) abandonQuest(msg *client.AbandonQuest)                           {}
func (p *player) shareQuest(msg *client.ShareQuest)                               {}
func (p *player) acceptReincarnation(msg *client.AcceptReincarnation)             {}
func (p *player) cancelReincarnation(msg *client.CancelReincarnation)             {}
func (p *player) combineItem(msg *client.CombineItem)                             {}
func (p *player) setConcentration(msg *client.SetConcentration)                   {}
func (p *player) awakeningNeedMaterials(msg *client.AwakeningNeedMaterials)       {}
func (p *player) awakeningLockedItem(msg *client.AwakeningLockedItem)             {}
func (p *player) awakening(msg *client.Awakening)                                 {}
func (p *player) disassembleItem(msg *client.DisassembleItem)                     {}
func (p *player) downgradeAwakening(msg *client.DowngradeAwakening)               {}
func (p *player) resetAddedItem(msg *client.ResetAddedItem)                       {}
func (p *player) sendMail(msg *client.SendMail)                                   {}
func (p *player) readMail(msg *client.ReadMail)                                   {}
func (p *player) collectParcel(msg *client.CollectParcel)                         {}
func (p *player) deleteMail(msg *client.DeleteMail)                               {}
func (p *player) lockMail(msg *client.LockMail)                                   {}
func (p *player) mailLockedItem(msg *client.MailLockedItem)                       {}
func (p *player) mailCost(msg *client.MailCost)                                   {}
func (p *player) updateIntelligentCreature(msg *client.UpdateIntelligentCreature) {}
func (p *player) intelligentCreaturePickup(msg *client.IntelligentCreaturePickup) {}
func (p *player) addFriend(msg *client.AddFriend)                                 {}
func (p *player) removeFriend(msg *client.RemoveFriend)                           {}
func (p *player) refreshFriends(msg *client.RefreshFriends)                       {}
func (p *player) addMemo(msg *client.AddMemo)                                     {}
func (p *player) guildBuffUpdate(msg *client.GuildBuffUpdate)                     {}
func (p *player) gameshopBuy(msg *client.GameshopBuy)                             {}
func (p *player) npcConfirmInput(msg *client.NPCConfirmInput)                     {}
func (p *player) reportIssue(msg *client.ReportIssue)                             {}
func (p *player) getRanking(msg *client.GetRanking)                               {}
func (p *player) opendoor(msg *client.Opendoor)                                   {}
func (p *player) getRentedItems(msg *client.GetRentedItems)                       {}
func (p *player) itemRentalRequest(msg *client.ItemRentalRequest)                 {}
func (p *player) itemRentalFee(msg *client.ItemRentalFee)                         {}
func (p *player) itemRentalPeriod(msg *client.ItemRentalPeriod)                   {}
func (p *player) depositRentalItem(msg *client.DepositRentalItem)                 {}
func (p *player) retrieveRentalItem(msg *client.RetrieveRentalItem)               {}
func (p *player) cancelItemRental(msg *client.CancelItemRental)                   {}
func (p *player) itemRentalLockFee(msg *client.ItemRentalLockFee)                 {}
func (p *player) itemRentalLockItem(msg *client.ItemRentalLockItem)               {}
func (p *player) confirmItemRental(msg *client.ConfirmItemRental)                 {}

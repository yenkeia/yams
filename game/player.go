package game

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/davyxu/cellnet"
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/client"
	"github.com/yenkeia/yams/game/proto/server"
)

type player struct {
	base
	session           *cellnet.Session
	actionList        *actionList
	gameStage         int
	accountID         int // account.ID
	characterID       int // character.ID 保存数据库用
	bindLocation      cm.Point
	bindMapID         int
	isDead            bool
	hp                int
	mp                int
	recovery          recovery
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
	allowTrade        bool
	sendedItemInfoIDs []int
	magics            map[cm.Spell]*userMagic
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
	accuracy          int
	agility           int
	criticalRate      int
	criticalDamage    int
	currentBagWeight  int
	maxBagWeight      int
	maxWearWeight     int
	maxHandWeight     int
	attackSpeed       int
	luck              int
	lifeOnHit         int
	hpDrainRate       int // hp 流失率
	reflect           int
	magicResist       int
	poisonResist      int
	healthRecovery    int
	manaRecovery      int
	poisonRecovery    int
	holy              int
	freezing          int
	poisonAttack      int
}

type recovery struct {
	// 生命药水回复
	hpPotValue    int           // 回复总值
	hpPotPerValue int           // 一次回复多少
	hpPotNextTime time.Time     // 下次生效时间
	hpPotDuration time.Duration // 两次生效时间间隔
	hpPotTickNum  int           // 总共跳几次
	hpPotTickTime int           // 当前第几跳
	// 魔法药水回复
	mpPotValue    int
	mpPotPerValue int
	mpPotNextTime time.Time
	mpPotDuration time.Duration
	mpPotTickNum  int
	mpPotTickTime int
	// 角色自身的生命/魔法回复
	recoveryNextTime time.Time
	recoveryDuration time.Duration
}

func (p *player) String() string {
	res := fmt.Sprintf(`
	玩家: %s, 等级: %d, 位置: %s, objectID: %d,
	当前血量 hp: %d, maxHP: %d, mp: %d, maxMP: %d
	当前经验值 experience: %d, maxExpericence: %d,
	生命值回复 healthRecovery: %d, 魔法值回复 manaRecovery: %d
	中毒回复？poisonRecovery: %d
	物理防御力minAC: %d, maxAC: %d
	魔法防御力minMAC: %d, maxMAC: %d
	攻击力minDC: %d, maxDC: %d
	魔法力minMC: %d, maxMC: %d
	道术力minSC: %d, maxSC: %d
	准确accuracy: %d
	敏捷agility: %d
	暴击率criticalRate: %d, 暴击伤害criticalDamage: %d
	currentBagWeight: %d, maxBagWeight: %d, maxWearWeight: %d, maxHandWeight: %d
	`, p.name, p.level, p.location, p.objectID,
		p.hp, p.maxHP, p.mp, p.maxMP,
		p.experience, p.maxExperience,
		p.healthRecovery, p.manaRecovery, p.poisonRecovery,
		p.minAC, p.maxAC, p.minMAC, p.maxMAC, p.minDC, p.maxDC, p.minMC, p.maxMC, p.minSC, p.maxSC,
		p.accuracy, p.agility, p.criticalRate, p.criticalDamage, p.currentBagWeight, p.maxBagWeight, p.maxWearWeight, p.maxHandWeight)
	return res
}

func (p *player) getObjectID() int {
	return p.objectID
}

func (p *player) getPosition() cm.Point {
	return p.location
}

func (p *player) isBlocking() bool {
	return !p.isDead
}

// TODO
func (p *player) isAttackTarget(attacker) bool {
	return true
}

// TODO
func (p *player) isFriendlyTarget(atk attacker) bool {
	return false
}

func (p *player) update(now time.Time) {
	p.actionList.execute(now)
	p.updateRecovery(now)
}

// 处理玩家自身回复，药水回复
func (p *player) updateRecovery(now time.Time) {
	rec := &p.recovery
	if rec.hpPotValue != 0 && rec.hpPotNextTime.Before(now) && p.hp != p.maxHP {
		p.changeHP(rec.hpPotPerValue)
		rec.hpPotTickTime++
		if rec.hpPotTickTime >= rec.hpPotTickNum {
			rec.hpPotValue = 0
		} else {
			rec.hpPotNextTime = now.Add(rec.hpPotDuration)
		}
	}
	if rec.mpPotValue != 0 && rec.mpPotNextTime.Before(now) && p.mp != p.maxMP {
		p.changeMP(rec.mpPotPerValue)
		rec.mpPotTickTime++
		if rec.mpPotTickTime >= rec.mpPotTickNum {
			rec.mpPotValue = 0
		} else {
			rec.mpPotNextTime = now.Add(rec.mpPotDuration)
		}
	}
	if now.After(rec.recoveryNextTime) {
		rec.recoveryNextTime = now.Add(rec.recoveryDuration)
		p.changeHP(int(float32(p.maxHP)*0.03) + 1 + p.healthRecovery)
		p.changeMP(int(float32(p.maxMP)*0.03) + 1 + p.manaRecovery)
	}
}

// 直接将 hp 设置为某个值
func (p *player) setHP(hp int) {
	if p.hp == hp {
		return
	}
	if hp <= 0 {
		hp = 0
	}
	if hp >= p.maxHP {
		hp = p.maxHP
	}
	p.hp = hp
	if !p.isDead && p.hp == 0 {
		p.die()
	}
	p.enqueue(&server.HealthChanged{HP: uint16(p.hp), MP: uint16(p.mp)})
	p.broadcastHealthChange()
	log.Debugf("setHP. hp: %d, p.hp: %d, p.maxHP: %d", hp, p.hp, p.maxHP)
}

func (p *player) setMP(mp int) {
	if p.mp == mp {
		return
	}
	if mp <= 0 {
		mp = 0
	}
	if mp >= p.maxMP {
		mp = p.maxMP
	}
	p.mp = mp
	p.enqueue(&server.HealthChanged{HP: uint16(p.hp), MP: uint16(p.mp)})
	p.broadcastHealthChange()
	log.Debugf("setMP. mp: %d, p.mp: %d, p.maxMP: %d", mp, p.mp, p.maxMP)
}

// 改变玩家血量 amount 可以是负数，表示扣血
func (p *player) changeHP(amount int) {
	// log.Debugf("changeHP. amount: %d", amount)
	if amount == 0 || p.isDead {
		return
	}
	hp := p.hp + amount
	if hp <= 0 {
		hp = 0
	}
	if hp >= p.maxHP {
		hp = p.maxHP
	}
	p.setHP(hp)
}

func (p *player) changeMP(amount int) {
	// log.Debugf("changeMP. amount: %d", amount)
	if amount == 0 || p.isDead {
		return
	}
	mp := p.mp + amount
	if mp <= 0 {
		mp = 0
	}
	if mp >= p.maxMP {
		mp = p.maxMP
	}
	p.setMP(mp)
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
	case *spell:
		p.enqueue(o.getSpellObjectInfo())
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

func (p *player) broadcastHealthChange() {
	percent := byte(float32(p.hp) / float32(p.maxHP) * 100)
	msg := &server.ObjectHealth{
		ObjectID: uint32(p.objectID),
		Percent:  percent,
		Expire:   5,
	}
	p.broadcast(msg)
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
	now := env.maps[p.mapID].now
	p.recovery = recovery{
		hpPotDuration:    1 * time.Second, // 药水回复
		mpPotDuration:    1 * time.Second,
		recoveryNextTime: now.Add(10 * time.Second), // 角色自身回复
		recoveryDuration: 10 * time.Second,
	}
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
	p.allowTrade = false
	p.sendedItemInfoIDs = make([]int, 0)
	p.magics = loadPlayerMagics(p.characterID)
}

func (p *player) getClientMagics() []*server.ClientMagic {
	res := make([]*server.ClientMagic, 0)
	for _, um := range p.magics {
		res = append(res, um.toServerClientMagic())
	}
	return res
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

func (p *player) refreshStats() {
	// FIXME 因为不像 C# 有对比最小值最大值，可能会有整型溢出问题
	log.Debugf("before refreshStats \n %s", p)
	p.refreshLevelStats()
	p.refreshBagWeight()
	p.refreshEquipmentStats()
	log.Debugf("after refreshStats \n %s", p)
}

func (p *player) refreshLevelStats() {
	baseStats := gdb.baseStatsMap[p.class]
	p.accuracy = baseStats.StartAccuracy
	p.agility = baseStats.StartAgility
	p.criticalRate = baseStats.StartCriticalRate
	p.criticalDamage = baseStats.StartCriticalDamage
	p.maxExperience = gdb.levelMaxExpMap[p.level]
	p.maxHP = 14 + int((float32(p.level)/baseStats.HPGain+baseStats.HPGainRate)*float32(p.level))
	// log.Debugln(fmt.Sprintf("=======> p.maxHP: %d, p.level: %d, baseStats.HPGain: %f, baseStats.HPGainRate: %f", p.maxHP, p.level, baseStats.HPGain, baseStats.HPGainRate))
	p.minAC = 0
	if baseStats.MinAC > 0 {
		p.minAC = p.level / baseStats.MinAC
	}
	p.maxAC = 0
	if baseStats.MaxAC > 0 {
		p.maxAC = p.level / baseStats.MaxAC
	}
	p.minMAC = 0
	if baseStats.MinMAC > 0 {
		p.minMAC = p.level / baseStats.MinMAC
	}
	p.maxMAC = 0
	if baseStats.MaxMAC > 0 {
		p.maxMAC = p.level / baseStats.MaxMAC
	}
	p.minDC = 0
	if baseStats.MinDC > 0 {
		p.minDC = p.level / baseStats.MinDC
	}
	p.maxDC = 0
	if baseStats.MaxDC > 0 {
		p.maxDC = p.level / baseStats.MaxDC
	}
	p.minMC = 0
	if baseStats.MinMC > 0 {
		p.minMC = p.level / baseStats.MinMC
	}
	p.maxMC = 0
	if baseStats.MaxMC > 0 {
		p.maxMC = p.level / baseStats.MaxMC
	}
	p.minSC = 0
	if baseStats.MinSC > 0 {
		p.minSC = p.level / baseStats.MinSC
	}
	p.maxSC = 0
	if baseStats.MaxSC > 0 {
		p.maxSC = p.level / baseStats.MaxSC
	}
	p.criticalRate = 0
	if baseStats.CritialRateGain > 0 {
		p.criticalRate = p.criticalRate + int(float32(p.level)/baseStats.CritialRateGain)
	}
	p.criticalDamage = 0
	if baseStats.CriticalDamageGain > 0 {
		p.criticalDamage = p.criticalDamage + int(float32(p.level)/baseStats.CriticalDamageGain)
	}
	p.maxBagWeight = 50 + int(float32(p.level)/baseStats.BagWeightGain*float32(p.level))
	p.maxWearWeight = 15 + int(float32(p.level)/baseStats.WearWeightGain*float32(p.level))
	p.maxHandWeight = 12 + int(float32(p.level)/baseStats.HandWeightGain*float32(p.level))
	switch p.class {
	case cm.MirClassWarrior:
		p.maxHP = 14 + int((float32(p.level)/baseStats.HPGain+baseStats.HPGainRate+float32(p.level)/20.0)*float32(p.level))
		p.maxMP = 11 + int((float32(p.level)*3.5)+(float32(p.level)*baseStats.MPGainRate))
	case cm.MirClassWizard:
		p.maxMP = 13 + int((float32(p.level/5.0+2.0)*2.2*float32(p.level))+(float32(p.level)*baseStats.MPGainRate))
	case cm.MirClassTaoist:
		p.maxMP = 13 + int((float32(p.level)/8.0*2.2*float32(p.level))+(float32(p.level)*baseStats.MPGainRate))
	}
}

func (p *player) refreshBagWeight() {
	p.currentBagWeight = 0
	for _, ui := range p.inventory.items {
		if ui != nil {
			p.currentBagWeight += ui.info.Weight
		}
	}
}

func (p *player) refreshEquipmentStats() {
	oldLooksWeapon := p.looksWeapon
	oldLooksWeaponEffect := p.looksWeaponEffect
	oldLooksArmour := p.looksArmour
	// oldMountType = MountType;
	oldLooksWings := p.looksWings
	oldLight := p.light

	p.looksArmour = 0
	p.looksWeapon = -1
	p.looksWeaponEffect = 0
	p.looksWings = 0

	hpRate := 0
	mpRate := 0
	acRate := 0
	macRate := 0

	for _, temp := range p.equipment.items {
		if temp == nil {
			continue
		}

		RealItem := gdb.getRealItem(temp.info, p.level, p.class, gdb.itemInfos)

		p.minAC = p.minAC + RealItem.MinAC
		p.maxAC = p.maxAC + RealItem.MaxAC + temp.ac
		p.minMAC = p.minMAC + RealItem.MinMAC
		p.maxMAC = p.maxMAC + RealItem.MaxMAC + temp.mac
		p.minDC = p.minDC + RealItem.MinDC
		p.maxDC = p.maxDC + RealItem.MaxDC + temp.dc
		p.minMC = p.minMC + RealItem.MinMC
		p.maxMC = p.maxMC + RealItem.MaxMC + temp.mc
		p.minSC = p.minSC + RealItem.MinSC
		p.maxSC = p.maxSC + RealItem.MaxSC + temp.sc
		p.maxHP = p.maxHP + RealItem.HP + temp.hp
		p.maxMP = p.maxMP + RealItem.MP + temp.mp

		p.maxBagWeight = p.maxBagWeight + RealItem.BagWeight
		p.maxWearWeight = p.maxWearWeight + RealItem.WearWeight
		p.maxHandWeight = p.maxHandWeight + RealItem.HandWeight

		p.attackSpeed = p.attackSpeed + temp.attackSpeed + RealItem.AttackSpeed
		p.luck = p.luck + temp.luck + RealItem.Luck

		p.accuracy = p.accuracy + RealItem.Accuracy + temp.accuracy
		p.agility = p.agility + RealItem.Agility + temp.agility

		hpRate = hpRate + RealItem.HpRate
		mpRate = mpRate + RealItem.MpRate
		acRate = acRate + RealItem.MaxAcRate
		macRate = macRate + RealItem.MaxMacRate

		p.magicResist = p.magicResist + temp.magicResist + RealItem.MagicResist
		p.poisonResist = p.poisonResist + temp.poisonResist + RealItem.PoisonResist
		p.healthRecovery = p.healthRecovery + temp.healthRecovery + RealItem.HealthRecovery
		p.manaRecovery = p.manaRecovery + temp.manaRecovery + RealItem.ManaRecovery
		p.poisonRecovery = p.poisonRecovery + temp.poisonRecovery + RealItem.PoisonRecovery
		p.criticalRate = p.criticalRate + temp.criticalRate + RealItem.CriticalRate
		p.criticalDamage = p.criticalDamage + temp.criticalDamage + RealItem.CriticalDamage
		p.holy = p.holy + RealItem.Holy
		p.freezing = p.freezing + temp.freezing + RealItem.Freezing
		p.poisonAttack = p.poisonAttack + temp.poisonAttack + RealItem.PoisonAttack
		p.reflect = p.reflect + RealItem.Reflect
		p.hpDrainRate = p.hpDrainRate + RealItem.HpDrainRate

		switch RealItem.Type {
		case cm.ItemTypeArmour:
			p.looksArmour = int(RealItem.Shape)
			p.looksWings = int(RealItem.Effect)
		case cm.ItemTypeWeapon:
			p.looksWeapon = int(RealItem.Shape)
			p.looksWeaponEffect = int(RealItem.Effect)
		}
	}

	p.maxHP = ((hpRate / 100) + 1) * p.maxHP
	p.maxMP = ((mpRate / 100) + 1) * p.maxMP
	p.maxAC = ((acRate / 100) + 1) * p.maxAC
	p.maxMAC = ((macRate / 100) + 1) * p.maxMAC

	/* TODO
	AddTempSkills(skillsToAdd);
	RemoveTempSkills(skillsToRemove);

	if (HasMuscleRing)
	{
		MaxBagWeight = (ushort)(MaxBagWeight * 2);
		MaxWearWeight = Math.Min(ushort.MaxValue, (ushort)(MaxWearWeight * 2));
		MaxHandWeight = Math.Min(ushort.MaxValue, (ushort)(MaxHandWeight * 2));
	}
	*/

	if (oldLooksArmour != p.looksArmour) || (oldLooksWeapon != p.looksWeapon) || (oldLooksWeaponEffect != p.looksWeaponEffect) || (oldLooksWings != p.looksWings) || (oldLight != p.light) {
		p.updateConcentration()
		p.broadcastPlayerUpdate()
	}
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

func (p *player) useItem(msg *client.UseItem) {
	var err error
	id := int(msg.UniqueID)
	res := &server.UseItem{UniqueID: uint64(msg.UniqueID), Success: false}
	index, item := p.getUserItemByObjectID(cm.MirGridTypeInventory, id)
	log.Debugf("index: %d, item: %s", index, item)
	if item == nil || index == -1 || p.isDead {
		p.enqueue(res)
		return
	}
	switch item.info.Type {
	case cm.ItemTypePotion:
		err = p.useItemPotion(item)
	case cm.ItemTypeScroll:
		err = p.useItemScroll(item)
	case cm.ItemTypeBook:
		err = p.giveSkill(cm.Spell(item.info.Shape), 1)
	case cm.ItemTypeScript:
		log.Errorln("不支持 item type script")
		p.enqueue(res)
		return
	}
	if err != nil {
		log.Errorf("player[%s] useItem error: %s", p.name, err)
	} else {
		if item.count > 1 {
			item.count--
		} else {
			err = p.inventory.set(index, nil)
		}
		res.Success = true
		p.receiveChat("恭喜你学会了新技能", cm.ChatTypeHint)
	}
	p.refreshBagWeight()
	p.enqueue(res)
}

// 使用药水
func (p *player) useItemPotion(ui *userItem) (err error) {
	switch ui.info.Shape {
	case 0: // NormalPotion 一般药水
		now := env.maps[p.mapID].now
		log.Debugf("useItemPotion, hpPotValue: %d, hpPotPerValue: %d", p.recovery.hpPotValue, p.recovery.hpPotPerValue)
		if ui.info.HP > 0 {
			p.recovery.hpPotValue = ui.info.HP                           // 回复总值
			p.recovery.hpPotPerValue = ui.info.HP / 3                    // 一次回复多少
			p.recovery.hpPotNextTime = now.Add(p.recovery.hpPotDuration) // 下次生效时间
			p.recovery.hpPotTickNum = 3                                  // 总共跳几次
			p.recovery.hpPotTickTime = 0                                 // 当前第几跳
		}
		if ui.info.MP > 0 {
			p.recovery.mpPotValue = ui.info.MP
			p.recovery.mpPotPerValue = ui.info.MP / 3
			p.recovery.mpPotNextTime = now.Add(p.recovery.mpPotDuration)
			p.recovery.mpPotTickNum = 3
			p.recovery.mpPotTickTime = 0
		}
		log.Debugf("useItemPotion, hpPotValue: %d, hpPotPerValue: %d", p.recovery.hpPotValue, p.recovery.hpPotPerValue)
	case 1: // SunPotion 太阳水
		p.changeHP(ui.info.HP)
		p.changeMP(ui.info.MP)
	default:
		return fmt.Errorf("不支持物品类型 userItem.info.Shape: %d", ui.info.Shape)
	}
	return nil
}

// TODO 使用卷轴
func (p *player) useItemScroll(ui *userItem) (err error) {
	return errors.New("暂不支持卷轴")
}

// 使用技能书
func (p *player) giveSkill(spell cm.Spell, level int) (err error) {
	info := gdb.spellMagicInfoMap[spell]
	if info == nil {
		return fmt.Errorf("没有找到技能, spell: %d", spell)
	}
	for _, um := range p.magics {
		if um.spell == spell {
			p.receiveChat("你已经学习该技能", cm.ChatTypeSystem)
			return fmt.Errorf("玩家[%s]已经学会该技能", p.name)
		}
	}
	um := newUserMagic(info, level, p.characterID, spell)
	um.id = pdb.addSkill(um)
	p.magics[um.spell] = um
	p.enqueue(&server.NewMagic{Magic: um.toServerClientMagic()})
	p.refreshStats()
	return nil
}

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
		item.broadcast(&server.ObjectRemove{ObjectID: uint32(item.objectID)})
	}
}

func (p *player) inspect(msg *client.Inspect) {
	id := int(msg.ObjectID)
	o := env.players[id]
	p.enqueue(&server.PlayerInspect{
		Name:      o.name,
		GuildName: o.guildName,
		GuildRank: o.guildRankName,
		Equipment: o.equipment.serverUserItems(),
		Class:     o.class,
		Gender:    o.gender,
		Hair:      uint8(o.hair),
		Level:     uint16(o.level),
		LoverName: "",
	})
}

func (p *player) changeAMode(msg *client.ChangeAMode) {
	p.attackMode = msg.Mode
	p.enqueue(&server.ChangeAMode{Mode: p.attackMode})
}

func (p *player) changePMode(msg *client.ChangePMode) {
	p.petMode = msg.Mode
	p.enqueue(&server.ChangePMode{Mode: p.petMode})
}

func (p *player) changeTrade(msg *client.ChangeTrade) {
	p.allowTrade = msg.AllowTrade
}

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
		obj := e.Value.(attackTarget)
		p.actionList.pushDelayAction(cm.DelayedTypeDamage, 300, func() { p.completeAttack(obj, damageFinal, defence, true) })
	}
}

func (p *player) completeAttack(args ...interface{}) {
	target := args[0].(attackTarget)
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
			p.levelMagic(magic)
			break
		}
	}
	*/
}

// winExp 根据怪物等级为玩家增加经验
func (p *player) winExp(amount, targetLevel int) {
	var expPoint int
	level := int(p.level)

	if level < targetLevel+10 { //|| !Settings.ExpMobLevelDifference
		expPoint = amount
	} else {
		expPoint = amount - int(math.Round(math.Max(float64(amount)/15.0, 1.0)*float64(level-(targetLevel+10))))
	}
	if expPoint <= 0 {
		expPoint = 1
	}
	log.Debugf("player winExp. expPoint: %d", expPoint)
	// if (GroupMembers != null)
	p.gainExp(expPoint)
}

func (p *player) gainExp(amount int) {
	p.experience += amount
	p.enqueue(&server.GainExperience{Amount: uint32(amount)})
	// log.Debugf("Player: %s GainExp amount: %d, p.Experience: %d, p.MaxExperience: %d\n", p.Name, amount, p.Experience, p.MaxExperience)
	if p.experience < p.maxExperience {
		return
	}
	// 连续升级
	var exp = p.experience
	for exp >= p.maxExperience {
		p.level++
		exp -= p.maxExperience
		p.refreshStats()
	}
	pdb.syncLevel(p)
	p.experience = exp
	p.levelUp()
}

// 玩家升级
func (p *player) levelUp() {
	p.refreshStats()
	p.setHP(p.maxHP)
	p.setMP(p.maxMP)
	p.enqueue(&server.LevelChanged{
		Level:         uint16(p.level),
		Experience:    int64(p.experience),
		MaxExperience: int64(p.maxExperience),
	})
	p.broadcast(&server.ObjectLeveled{ObjectID: uint32(p.objectID)})
}

// TODO 被攻击
func (p *player) attacked(atk attacker, dmg int, typ cm.DefenceType, isWeapon bool) int {
	return 0
}

// 玩家角色死亡
func (p *player) die() {
	p.hp = 0
	p.isDead = true
	p.enqueue(&server.Death{Direction: p.direction, LocationX: int32(p.location.X), LocationY: int32(p.location.Y)})
	p.broadcast(&server.ObjectDied{ObjectID: uint32(p.objectID), Direction: p.direction, LocationX: int32(p.location.X), LocationY: int32(p.location.Y), Type: 0})
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

func (p *player) craftItem(msg *client.CraftItem)     {}
func (p *player) sellItem(msg *client.SellItem)       {}
func (p *player) repairItem(msg *client.RepairItem)   {}
func (p *player) buyItemBack(msg *client.BuyItemBack) {}
func (p *player) sRepairItem(msg *client.SRepairItem) {}

func (p *player) magicKey(msg *client.MagicKey) {
	key := int(msg.Key)
	for _, um := range p.magics {
		if um.spell == msg.Spell {
			um.key = key
			pdb.syncMagicKey(p.characterID, um.spell, key)
		} else if um.key == int(msg.Key) {
			um.key = 0
			pdb.syncMagicKey(p.characterID, um.spell, 0)
		}
	}
}

func (p *player) magic(msg *client.Magic) {
	var (
		ok             bool
		err            error
		um             *userMagic
		cast           bool
		target         attackTarget
		spell          cm.Spell
		targetLocation cm.Point
	)
	spell = msg.Spell

	// if !p.canCast() {
	um, ok = p.magics[spell]
	if !ok {
		p.enqueue(&server.UserLocation{Location: p.location, Direction: p.direction})
		return
	}
	info := gdb.spellMagicInfoMap[spell]
	cost := info.BaseCost + info.LevelCost*um.level
	if cost > p.mp {
		p.enqueue(&server.UserLocation{Location: p.location, Direction: p.direction})
		return
	}
	p.direction = msg.Direction
	p.changeMP(-cost)
	targetID := int(msg.TargetID)
	env.maps[p.mapID].rangeObject(msg.Location, 1, func(o mapObject) bool {
		if o.getObjectID() == targetID {
			target = o.(attackTarget)
			return false
		}
		return true
	})
	targetLocation = msg.Location
	if target != nil {
		targetLocation = target.getPosition()
	}
	ctx := &magicContext{
		player:      p,
		spell:       spell,
		target:      target,
		targetPoint: msg.Location,
	}
	// cast 代表是否释放成功
	// targetID 上面 msg.TargetID 只是用来获取攻击目标 id
	// 		这里返回的 targetID 如果是 0 代表未击中目标，返回 msg.TargetID 表示击中目标造成伤害
	// err 是提示打印日志用的，和游戏逻辑无关
	targetID, err = startMagic(ctx)
	cast = true
	if err != nil {
		cast = false
		p.receiveChat(err.Error(), cm.ChatTypeSystem)
	}
	p.enqueue(&server.UserLocation{Location: p.location, Direction: p.direction})
	p.enqueue(&server.Magic{
		Spell:    msg.Spell,
		TargetID: uint32(targetID),
		TargetX:  int32(targetLocation.X),
		TargetY:  int32(targetLocation.Y),
		Cast:     cast,
		Level:    uint8(um.level),
	})
	p.broadcast(&server.ObjectMagic{
		ObjectID:      uint32(p.objectID),
		LocationX:     int32(p.location.X),
		LocationY:     int32(p.location.Y),
		Direction:     p.direction,
		Spell:         msg.Spell,
		TargetID:      uint32(targetID),
		TargetX:       int32(targetLocation.X),
		TargetY:       int32(targetLocation.Y),
		Cast:          cast,
		Level:         uint8(um.level),
		SelfBroadcast: false,
	})
}

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

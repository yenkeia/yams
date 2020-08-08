package game

import (
	"fmt"

	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

// userItem 是 yams 内部使用的
// orm.UserItem 是存在数据库里面的
// server.UserItem 是只用来和客户端交互用的
type userItem struct {
	info           *orm.ItemInfo
	id             int // 保存 orm.UserItem.ID
	objectID       int
	currentDura    int
	maxDura        int
	count          int
	soulBoundID    int
	hp             int // 相对于 orm.ItemInfo 额外增加的属性，也就是所谓的小极品
	mp             int
	ac             int
	mac            int
	dc             int
	mc             int
	sc             int
	accuracy       int
	agility        int
	attackSpeed    int
	luck           int
	magicResist    int
	poisonResist   int
	healthRecovery int
	manaRecovery   int
	poisonRecovery int
	hpRate         int
	mpRate         int
	criticalRate   int
	criticalDamage int
	holy           int
	freezing       int
	poisonAttack   int
	bools          int
	strong         int
}

func (ui *userItem) String() string {
	return fmt.Sprintf("userItem: %s, objectID: %d", ui.info.Name, ui.objectID)
}

// newUserItem 新建一个游戏内部使用的 userItem
func newUserItem(info *orm.ItemInfo) *userItem {
	return &userItem{
		info:        info,
		objectID:    env.newObjectID(),
		count:       1,
		soulBoundID: -1, // 根据客户端代码 -1 时候不现实 "灵魂绑定给" ..
	}
}

// newUserItemFromORM 从数据库加载 userItem
func newUserItemFromORM(ui *orm.UserItem) *userItem {
	return &userItem{
		info:        gdb.itemInfoMap[ui.ItemID],
		id:          ui.ID,
		objectID:    env.newObjectID(),
		currentDura: ui.CurrentDura,
		maxDura:     ui.MaxDura,
		count:       ui.Count,
		soulBoundID: ui.SoulBoundID,
	}
}

// serverUserItem 把内部使用的 userItem 转换成和客户端沟通使用的 server.UserItem
func (ui *userItem) serverUserItem() *server.UserItem {
	return &server.UserItem{
		ID:             uint64(ui.objectID),
		ItemID:         int32(ui.info.ID), // itemInfo.ID
		CurrentDura:    uint16(ui.currentDura),
		MaxDura:        uint16(ui.maxDura),
		Count:          uint32(ui.count),
		AC:             uint8(ui.ac),
		MAC:            uint8(ui.mac),
		DC:             uint8(ui.dc),
		MC:             uint8(ui.mc),
		SC:             uint8(ui.sc),
		Accuracy:       uint8(ui.accuracy),
		Agility:        uint8(ui.agility),
		HP:             uint8(ui.hp),
		MP:             uint8(ui.mp),
		AttackSpeed:    int8(ui.attackSpeed),
		Luck:           int8(ui.luck),
		SoulBoundId:    uint32(ui.soulBoundID), // uint32(ui.SoulBoundID),
		Bools:          uint8(ui.bools),
		Strong:         uint8(ui.strong),
		MagicResist:    uint8(ui.magicResist),
		PoisonResist:   uint8(ui.poisonResist),
		HealthRecovery: uint8(ui.healthRecovery),
		ManaRecovery:   uint8(ui.manaRecovery),
		PoisonRecovery: uint8(ui.poisonRecovery),
		CriticalRate:   uint8(ui.criticalRate),
		CriticalDamage: uint8(ui.criticalDamage),
		Freezing:       uint8(ui.freezing),
		PoisonAttack:   uint8(ui.poisonAttack),
	}
}

func (ui *userItem) ormUserItem() *orm.UserItem {
	return &orm.UserItem{
		ID:             ui.id,
		ItemID:         ui.info.ID,
		CurrentDura:    ui.currentDura,
		MaxDura:        ui.maxDura,
		Count:          ui.count,
		AC:             ui.ac,
		MAC:            ui.mac,
		DC:             ui.dc,
		MC:             ui.mc,
		SC:             ui.sc,
		Accuracy:       ui.accuracy,
		Agility:        ui.agility,
		HP:             ui.hp,
		MP:             ui.mp,
		AttackSpeed:    ui.attackSpeed,
		Luck:           ui.luck,
		SoulBoundID:    ui.soulBoundID,
		Bools:          ui.bools,
		Strong:         ui.strong,
		MagicResist:    ui.magicResist,
		PoisonResist:   ui.poisonResist,
		HealthRecovery: ui.healthRecovery,
		ManaRecovery:   ui.manaRecovery,
		PoisonRecovery: ui.poisonRecovery,
		CriticalRate:   ui.criticalRate,
		CriticalDamage: ui.criticalDamage,
		Freezing:       ui.freezing,
		PoisonAttack:   ui.poisonAttack,
	}
}

// TODO
func (ui *userItem) price() int {
	return ui.info.Price
}

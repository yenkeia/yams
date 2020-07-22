package game

import (
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

// userItem 是 yams 内部使用的
// orm.UserItem 是存在数据库里面的
// server.UserItem 是只用来和客户端交互用的
type userItem struct {
	info        *orm.ItemInfo
	id          int // 保存 orm.UserItem.ID
	objectID    int
	currentDura int
	maxDura     int
	count       int
	soulBoundID int
	ac          int // 额外增加的属性
	mac         int // 额外增加的属性
	dc          int // 额外增加的属性
	mc          int // 额外增加的属性
	sc          int // 额外增加的属性
	accuracy    int // 额外增加的属性
	agility     int // 额外增加的属性
	attackSpeed int // 额外增加的属性
	luck        int // 额外增加的属性
}

// newUserItem 新建一个游戏内部使用的 userItem
func newUserItem(info *orm.ItemInfo) *userItem {
	return &userItem{
		info:     info,
		objectID: env.newObjectID(),
		count:    1,
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
		HP:             0, // uint8(ui.HP),
		MP:             0, // uint8(ui.MP),
		AttackSpeed:    int8(ui.attackSpeed),
		Luck:           int8(ui.luck),
		SoulBoundId:    uint32(ui.soulBoundID), // uint32(ui.SoulBoundID),
		Bools:          0,                      // uint8(ui.Bools),
		Strong:         0,                      // uint8(ui.Strong),
		MagicResist:    0,                      // uint8(ui.MagicResist),
		PoisonResist:   0,                      // uint8(ui.PoisonResist),
		HealthRecovery: 0,                      // uint8(ui.HealthRecovery),
		ManaRecovery:   0,                      // uint8(ui.ManaRecovery),
		PoisonRecovery: 0,                      // uint8(ui.PoisonRecovery),
		CriticalRate:   0,                      // uint8(ui.CriticalRate),
		CriticalDamage: 0,                      // uint8(ui.CriticalDamage),
		Freezing:       0,                      // uint8(ui.Freezing),
		PoisonAttack:   0,                      // uint8(ui.PoisonAttack),
	}
}

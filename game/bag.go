package game

import (
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

type bag struct {
	items []*orm.UserItem
}

func (b *bag) convertItems() []*server.UserItem {
	res := make([]*server.UserItem, len(b.items))
	for i := range b.items {
		item := b.items[i]
		if item == nil {
			res[i] = nil
			continue
		}
		res[i] = &server.UserItem{
			ID:             uint64(item.ID),            // uint64 `gorm:"primary_key"` // UniqueID
			ItemID:         int32(item.ItemID),         // int32
			CurrentDura:    uint16(item.CurrentDura),   // uint16
			MaxDura:        uint16(item.MaxDura),       // uint16
			Count:          uint32(item.Count),         // uint32
			AC:             uint8(item.AC),             // uint8
			MAC:            uint8(item.MAC),            // uint8
			DC:             uint8(item.DC),             // uint8
			MC:             uint8(item.MC),             // uint8
			SC:             uint8(item.SC),             // uint8
			Accuracy:       uint8(item.Accuracy),       // uint8
			Agility:        uint8(item.Agility),        // uint8
			HP:             uint8(item.HP),             // uint8
			MP:             uint8(item.MP),             // uint8
			AttackSpeed:    int8(item.AttackSpeed),     // int8
			Luck:           int8(item.Luck),            // int8
			SoulBoundId:    uint32(item.SoulBoundID),   // uint32
			Bools:          uint8(item.Bools),          // uint8
			Strong:         uint8(item.Strong),         // uint8
			MagicResist:    uint8(item.MagicResist),    // uint8
			PoisonResist:   uint8(item.PoisonResist),   // uint8
			HealthRecovery: uint8(item.HealthRecovery), // uint8
			ManaRecovery:   uint8(item.ManaRecovery),   // uint8
			PoisonRecovery: uint8(item.PoisonRecovery), // uint8
			CriticalRate:   uint8(item.CriticalRate),   // uint8
			CriticalDamage: uint8(item.CriticalDamage), // uint8
			Freezing:       uint8(item.Freezing),       // uint8
			PoisonAttack:   uint8(item.PoisonAttack),   // uint8
		}
	}
	return res
}

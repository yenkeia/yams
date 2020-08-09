package game

import (
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

// userMagic 是玩家已经学会的的技能
type userMagic struct {
	info        *orm.MagicInfo
	id          int // orm.UserMagic.ID
	characterID int
	magicID     int // info.ID
	spell       cm.Spell
	level       int
	key         int
	experience  int
	isTempSpell bool
	castTime    int
}

func newUserMagic(info *orm.MagicInfo, level int, characterID int, spell cm.Spell) *userMagic {
	return &userMagic{
		info:        gdb.spellMagicInfoMap[spell],
		characterID: characterID,
		magicID:     info.ID, // int // info.ID
		spell:       spell,   // cm.Spell
		level:       level,   // int
		key:         0,       // int
		experience:  0,       // int
		isTempSpell: false,   // bool
		castTime:    0,       // int
	}
}

func newUserMagicFromORM(um *orm.UserMagic) *userMagic {
	return &userMagic{
		info:        gdb.spellMagicInfoMap[cm.Spell(um.Spell)],
		id:          um.ID,
		characterID: um.CharacterID,
		magicID:     um.MagicID,         // int // info.ID
		spell:       cm.Spell(um.Spell), // cm.Spell
		level:       um.Level,           // int
		key:         um.Key,             // int
		experience:  um.Experience,      // int
		isTempSpell: um.IsTempSpell,     // bool
		castTime:    um.CastTime,        // int
	}
}

func (um *userMagic) ormUserMagic() *orm.UserMagic {
	return &orm.UserMagic{
		// ID          int `gorm:"primary_key"`
		CharacterID: um.characterID,
		MagicID:     um.magicID,
		Spell:       int(um.spell),
		Level:       um.level,
		Key:         um.key,
		Experience:  um.experience,
		IsTempSpell: um.isTempSpell,
		CastTime:    um.castTime,
	}
}

func (um *userMagic) toServerClientMagic() *server.ClientMagic {
	//castTime := (CastTime != 0) && (SMain.Envir.Time > CastTime) ? SMain.Envir.Time - CastTime : 0
	delay := um.info.DelayBase - (um.level * um.info.DelayReduction)
	castTime := 0
	return &server.ClientMagic{
		Spell:      um.spell,
		BaseCost:   uint8(um.info.BaseCost),
		LevelCost:  uint8(um.info.LevelCost),
		Icon:       uint8(um.info.Icon),
		Level1:     uint8(um.info.Level1),
		Level2:     uint8(um.info.Level2),
		Level3:     uint8(um.info.Level3),
		Need1:      uint16(um.info.Need1),
		Need2:      uint16(um.info.Need2),
		Need3:      uint16(um.info.Need3),
		Level:      uint8(um.level),
		Key:        uint8(um.key),
		Experience: uint16(um.experience),
		Delay:      int64(delay),
		Range:      uint8(um.info.MagicRange),
		CastTime:   int64(castTime),
	}
}

func loadPlayerMagics(characterID int) map[cm.Spell]*userMagic {
	res := make(map[cm.Spell]*userMagic)
	magics := make([]*orm.UserMagic, 0)
	pdb.db.Table("user_magic").Where("character_id = ?", characterID).Find(&magics)
	for _, m := range magics {
		res[cm.Spell(m.Spell)] = newUserMagicFromORM(m)
	}
	return res
}

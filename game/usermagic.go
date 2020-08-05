package game

import (
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
)

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

func newUserMagicFromORM(um *orm.UserMagic) *userMagic {
	return &userMagic{
		info:        gdb.magicInfoMap[um.MagicID],
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

func loadPlayerMagics(characterID int) []*userMagic {
	res := make([]*userMagic, 0)
	magics := make([]*orm.UserMagic, 0)
	pdb.db.Table("user_magic").Where("character_id = ?", characterID).Find(&magics)
	for _, m := range magics {
		res = append(res, newUserMagicFromORM(m))
	}
	return res
}

package game

import (
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
)

type userMagic struct {
	info        *orm.MagicInfo
	id          int `gorm:"primary_key"`
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
	return nil
}

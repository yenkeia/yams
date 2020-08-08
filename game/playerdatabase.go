package game

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
)

type playerDatabase struct {
	db *gorm.DB
}

func newPlayerDatabase() *playerDatabase {
	name := conf.Mysql.PlayerDatabase
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, name))
	// defer accountDB.Close()	Close 之后对数据库的操作无效且不报错..
	if err != nil {
		panic(err)
	}
	pdb = new(playerDatabase)
	pdb.db = db
	return pdb
}

func (d *playerDatabase) setCharacterAttr(p *player, attr string, value interface{}) {
	d.db.Table("character").Where("id = ?", p.characterID).Update(attr, value)
}

func (d *playerDatabase) syncPosition(p *player) {
	d.setCharacterAttr(p, "current_map_id", p.mapID)
	d.setCharacterAttr(p, "direction", p.direction)
	d.setCharacterAttr(p, "current_location_x", p.location.X)
	d.setCharacterAttr(p, "current_location_y", p.location.Y)
}

func (d *playerDatabase) syncGold(p *player) {
	d.setCharacterAttr(p, "gold", p.gold)
}

func (d *playerDatabase) syncLevel(p *player) {
	d.setCharacterAttr(p, "level", p.level)
}

func (d *playerDatabase) syncExperience(p *player) {
	d.setCharacterAttr(p, "experience", p.experience)
}

func (d *playerDatabase) syncHPMP(p *player) {
	d.setCharacterAttr(p, "hp", p.hp)
	d.setCharacterAttr(p, "mp", p.mp)
}

func (d *playerDatabase) addSkill(um *userMagic) int {
	tmp := &orm.UserMagic{
		CharacterID: um.characterID,
		MagicID:     um.magicID,
		Spell:       um.info.Spell,
		Level:       um.level,
		Key:         um.key,
		Experience:  um.experience,
		IsTempSpell: um.isTempSpell,
		CastTime:    um.castTime,
	}
	d.db.Table("user_magic").Create(tmp)
	return tmp.ID
}

func (d *playerDatabase) syncMagicKey(characterID int, spell cm.Spell, key int) {
	d.db.Table("user_magic").Where("character_id = ? and spell = ?", characterID, spell).Update("magic_key", key)
}

package game

import (
	"fmt"

	"github.com/jinzhu/gorm"
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

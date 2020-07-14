package game

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yenkeia/yams/game/orm"
)

type mirData struct {
	mapInfos []*orm.MapInfo
	npcInfos []*orm.NPCInfo
}

func newmirData() *mirData {
	name := conf.Mysql.DataDB
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, name))
	defer db.Close()
	if err != nil {
		panic(err)
	}
	mirData := new(mirData)
	db.Table("map_info").Find(&mirData.mapInfos)
	db.Table("npc_info").Find(&mirData.npcInfos)
	return mirData
}

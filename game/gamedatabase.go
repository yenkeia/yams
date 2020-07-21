package game

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yenkeia/yams/game/orm"
)

type gameDatabase struct {
	mapInfos     []*orm.MapInfo
	npcInfos     []*orm.NPCInfo
	itemInfos    []*orm.ItemInfo
	monsterInfos []*orm.MonsterInfo
	respawnInfos []*orm.RespawnInfo
}

func newGameDatabase() *gameDatabase {
	name := conf.Mysql.GameDatabase
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, name))
	defer db.Close()
	if err != nil {
		panic(err)
	}
	gameData := new(gameDatabase)
	db.Table("map_info").Find(&gameData.mapInfos)
	db.Table("npc_info").Find(&gameData.npcInfos)
	db.Table("item_info").Find(&gameData.itemInfos)
	db.Table("monster_info").Find(&gameData.monsterInfos)
	db.Table("respawn_info").Find(&gameData.respawnInfos)
	return gameData
}

func (data *gameDatabase) getMonsterInfo(id int) *orm.MonsterInfo {
	for _, mi := range data.monsterInfos {
		if mi.ID == id {
			return mi
		}
	}
	return nil
}

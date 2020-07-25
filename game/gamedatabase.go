package game

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yenkeia/yams/game/orm"
)

type gameDatabase struct {
	mapInfos        []*orm.MapInfo
	npcInfos        []*orm.NPCInfo
	itemInfos       []*orm.ItemInfo
	monsterInfos    []*orm.MonsterInfo
	respawnInfos    []*orm.RespawnInfo
	itemInfoMap     map[int]*orm.ItemInfo    // key: orm.ItemInfo.ID
	itemInfoNameMap map[string]*orm.ItemInfo // key: orm.ItemInfo.Name
	monsterInfoMap  map[int]*orm.MonsterInfo // key: orm.MonsterInfo.ID
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
	gameData.itemInfoMap = make(map[int]*orm.ItemInfo)
	gameData.monsterInfoMap = make(map[int]*orm.MonsterInfo)
	gameData.itemInfoNameMap = make(map[string]*orm.ItemInfo)
	for _, ii := range gameData.itemInfos {
		gameData.itemInfoMap[ii.ID] = ii
		gameData.itemInfoNameMap[ii.Name] = ii
	}
	for _, mi := range gameData.monsterInfos {
		gameData.monsterInfoMap[mi.ID] = mi
	}
	return gameData
}

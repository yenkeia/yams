package game

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
)

type gameDatabase struct {
	mapInfos        []*orm.MapInfo
	npcInfos        []*orm.NPCInfo
	itemInfos       []*orm.ItemInfo
	monsterInfos    []*orm.MonsterInfo
	respawnInfos    []*orm.RespawnInfo
	baseStats       []*orm.BaseStats
	itemInfoMap     map[int]*orm.ItemInfo          // key: orm.ItemInfo.ID
	itemInfoNameMap map[string]*orm.ItemInfo       // key: orm.ItemInfo.Name
	monsterInfoMap  map[int]*orm.MonsterInfo       // key: orm.MonsterInfo.ID
	baseStatsMap    map[cm.MirClass]*orm.BaseStats // 各职业基础属性
	levelMaxExpMap  map[int]int                    // 玩家等级和最大经验对应关系
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
	db.Table("base_stats").Find(&gameData.baseStats)
	gameData.itemInfoMap = make(map[int]*orm.ItemInfo)
	gameData.monsterInfoMap = make(map[int]*orm.MonsterInfo)
	gameData.itemInfoNameMap = make(map[string]*orm.ItemInfo)
	gameData.baseStatsMap = make(map[cm.MirClass]*orm.BaseStats)
	gameData.levelMaxExpMap = make(map[int]int)
	for _, ii := range gameData.itemInfos {
		gameData.itemInfoMap[ii.ID] = ii
		gameData.itemInfoNameMap[ii.Name] = ii
	}
	for _, mi := range gameData.monsterInfos {
		gameData.monsterInfoMap[mi.ID] = mi
	}
	for _, s := range gameData.baseStats {
		gameData.baseStatsMap[cm.MirClass(s.ID-1)] = s
	}
	var levelExperience []*orm.LevelMaxExperience
	db.Table("level_max_experience").Find(&levelExperience)
	for _, i := range levelExperience {
		gameData.levelMaxExpMap[i.ID] = i.MaxExperience
	}
	return gameData
}

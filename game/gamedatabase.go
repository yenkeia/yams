package game

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
)

type gameDatabase struct {
	mapInfos          []*orm.MapInfo
	npcInfos          []*orm.NPCInfo
	itemInfos         []*orm.ItemInfo
	monsterInfos      []*orm.MonsterInfo
	respawnInfos      []*orm.RespawnInfo
	baseStats         []*orm.BaseStats
	magicInfos        []*orm.MagicInfo
	itemInfoMap       map[int]*orm.ItemInfo          // key: orm.ItemInfo.ID
	itemInfoNameMap   map[string]*orm.ItemInfo       // key: orm.ItemInfo.Name
	monsterInfoMap    map[int]*orm.MonsterInfo       // key: orm.MonsterInfo.ID
	baseStatsMap      map[cm.MirClass]*orm.BaseStats // 各职业基础属性
	levelMaxExpMap    map[int]int                    // 玩家等级和最大经验对应关系
	spellMagicInfoMap map[cm.Spell]*orm.MagicInfo    // key: orm.MagicInfo.Spell
	dropInfoMap       map[string][]*dropInfo         // key: orm.MonsterInfo.Name, value: 怪物掉落物品列表
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
	db.Table("magic_info").Find(&gameData.magicInfos)
	gameData.itemInfoMap = make(map[int]*orm.ItemInfo)
	gameData.monsterInfoMap = make(map[int]*orm.MonsterInfo)
	gameData.itemInfoNameMap = make(map[string]*orm.ItemInfo)
	gameData.baseStatsMap = make(map[cm.MirClass]*orm.BaseStats)
	gameData.levelMaxExpMap = make(map[int]int)
	gameData.spellMagicInfoMap = make(map[cm.Spell]*orm.MagicInfo)
	gameData.dropInfoMap = make(map[string][]*dropInfo)
	for _, ii := range gameData.itemInfos {
		gameData.itemInfoMap[ii.ID] = ii
		gameData.itemInfoNameMap[ii.Name] = ii
		b1 := ii.Bools & 0x04
		if b1 == 0x04 {
			ii.ClassBased = true
		}
		b2 := ii.Bools & 0x08
		if b2 == 0x08 {
			ii.LevelBased = true
		}
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
	for _, m := range gameData.magicInfos {
		gameData.spellMagicInfoMap[cm.Spell(m.Spell)] = m
	}
	// 怪物物品掉落
	dropDir := conf.Assets + "/Drops/"
	files, err := ioutil.ReadDir(dropDir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		dropInfos, err := loadDropFile(dropDir+file.Name(), gameData.itemInfoNameMap)
		if err != nil {
			panic(err)
		}
		tmp := []rune(file.Name())
		name := string(tmp[:len(tmp)-4])
		gameData.dropInfoMap[name] = dropInfos
	}
	return gameData
}

func (d *gameDatabase) getRealItem(origin *orm.ItemInfo, level int, job cm.MirClass, itemList []*orm.ItemInfo) *orm.ItemInfo {
	if origin.ClassBased && origin.LevelBased {
		return d.getClassAndLevelBasedItem(origin, job, level, itemList)
	}
	if origin.ClassBased {
		return d.getClassBasedItem(origin, job, itemList)
	}
	if origin.LevelBased {
		return d.getLevelBasedItem(origin, level, itemList)
	}
	return origin
}

func (d *gameDatabase) getLevelBasedItem(origin *orm.ItemInfo, level int, itemList []*orm.ItemInfo) *orm.ItemInfo {
	output := origin
	for i := 0; i < len(itemList); i++ {
		info := itemList[i]
		// if info.Name.StartsWith(Origin.Name) {
		if strings.HasPrefix(info.Name, origin.Name) {
			if (info.RequiredType == cm.RequiredTypeLevel) && (int(info.RequiredAmount) <= level) && (output.RequiredAmount < info.RequiredAmount) && (origin.RequiredGender == info.RequiredGender) {
				output = info
			}
		}
	}
	return output
}

func (d *gameDatabase) getClassBasedItem(origin *orm.ItemInfo, job cm.MirClass, itemList []*orm.ItemInfo) *orm.ItemInfo {
	for i := 0; i < len(itemList); i++ {
		info := itemList[i]
		if strings.HasPrefix(info.Name, origin.Name) {
			if (uint8(info.RequiredClass) == (1 << uint8(job))) && (origin.RequiredGender == info.RequiredGender) {
				return info
			}
		}
	}
	return origin
}

func (d *gameDatabase) getClassAndLevelBasedItem(origin *orm.ItemInfo, job cm.MirClass, level int, itemList []*orm.ItemInfo) *orm.ItemInfo {
	output := origin
	for i := 0; i < len(itemList); i++ {
		info := itemList[i]
		if strings.HasPrefix(info.Name, origin.Name) {
			if uint8(info.RequiredClass) == (1 << uint8(job)) {
				if (info.RequiredType == cm.RequiredTypeLevel) && (int(info.RequiredAmount) <= level) && (output.RequiredAmount <= info.RequiredAmount) && (origin.RequiredGender == info.RequiredGender) {
					output = info
				}
			}
		}
	}
	return output
}

package game

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yenkeia/yams/game/orm"
)

type gameData struct {
	mapInfos []*orm.MapInfo
}

func newGameData() *gameData {
	name := "gamedata"
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, name))
	defer db.Close()
	if err != nil {
		panic(err)
	}
	gameData := new(gameData)
	db.Table("map_info").Find(&gameData.mapInfos)
	return gameData
}

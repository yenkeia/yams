package game

import "github.com/yenkeia/yams/game/orm"

type gameData struct {
	mapInfos []*orm.MapInfo
}

func newGameData() *gameData {
	return nil
}

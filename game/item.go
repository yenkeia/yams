package game

import (
	"github.com/yenkeia/yams/game/cm"
)

// item 是地图上显示的游戏物品
type item struct {
	baseObject
	gold int
	ui   *userItem
}

// TODO
func newItem() *item {
	return nil
}

// TODO
func newItemGold() *item {
	return nil
}

func (i *item) getObjectID() int {
	return i.objectID
}

func (i *item) getPosition() cm.Point {
	return i.location
}

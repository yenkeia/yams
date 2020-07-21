package game

import (
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
)

type item struct {
	baseObject
	gold     int
	userItem *orm.UserItem
}

// TODO
func newItem() *item {
	return nil
}

// TODO
func newItemGold() *item {
	return nil
}

// TODO
func newUserItem() *orm.UserItem {
	return nil
}

func (i *item) getObjectID() int {
	return i.objectID
}

func (i *item) getPosition() cm.Point {
	return i.location
}

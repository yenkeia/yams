package game

import "github.com/yenkeia/yams/game/cm"

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

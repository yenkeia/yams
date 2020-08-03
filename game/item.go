package game

import (
	"fmt"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/proto/server"
)

// item 是地图上显示的游戏物品
type item struct {
	base
	gold int
	ui   *userItem
}

func newItem(mapID int, location cm.Point, ui *userItem) *item {
	i := new(item)
	i.objectID = env.newObjectID()
	i.name = ui.info.Name
	i.nameColor = cm.ColorWhite
	i.mapID = mapID
	i.location = location
	i.ui = ui
	return i
}

func newItemGold(mapID int, location cm.Point, gold int) *item {
	i := new(item)
	i.objectID = env.newObjectID()
	i.mapID = mapID
	i.location = location
	i.gold = gold
	return i
}

func (i *item) getObjectID() int {
	return i.objectID
}

func (i *item) getPosition() cm.Point {
	return i.location
}

func (i *item) drop(center cm.Point, distance int) (err error) {
	ok := false
	mp := env.maps[i.mapID]
	mp.rangeCell(center, distance, func(c *cell, x, y int) bool {
		if c == nil || c.hasItem() {
			return true
		}
		ok = true
		i.location = cm.NewPoint(x, y)
		mp.addObject(i)
		env.items[i.objectID] = i
		i.broadcast(i.getItemObjectInfo())
		return false
	})
	if !ok {
		return fmt.Errorf("坐标(%s)附近没有合适的点放置物品", center)
	}
	return nil
}

func (i *item) broadcast(msg interface{}) {
	mp := env.maps[i.mapID]
	mp.broadcast(i.location, msg, i.objectID)
}

func (i *item) getItemObjectInfo() interface{} {
	if i.ui == nil {
		return &server.ObjectGold{
			ObjectID:  uint32(i.objectID),  // uint32
			Gold:      uint32(i.gold),      // uint32
			LocationX: int32(i.location.X), // int32
			LocationY: int32(i.location.Y), // int32
		}
	}
	return &server.ObjectItem{
		ObjectID:  uint32(i.objectID),      // uint32
		Name:      i.name,                  // string
		NameColor: cm.ColorWhite.ToInt32(), // int32
		LocationX: int32(i.location.X),     // int32
		LocationY: int32(i.location.Y),     // int32
		Image:     uint16(i.ui.info.Image), // uint16
		Grade:     cm.ItemGradeNone,        // TODO cm.ItemGrade
	}
}

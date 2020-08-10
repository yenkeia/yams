package game

import (
	"container/list"

	"github.com/yenkeia/yams/game/cm"
)

type cell struct {
	attribute cm.CellAttribute
	objects   *list.List
}

func (c *cell) addObject(obj mapObject) {
	c.objects.PushBack(obj)
}

func (c *cell) contains(obj mapObject) (bool, *list.Element) {
	for e := c.objects.Front(); e != nil; e = e.Next() {
		if e.Value == obj {
			return true, e
		}
	}
	return false, nil
}

func (c *cell) deleteObject(obj mapObject) {
	if contain, e := c.contains(obj); contain {
		c.objects.Remove(e)
	}
}

// 用来判断地图点是否能摆放物品
func (c *cell) hasItem() bool {
	for it := c.objects.Front(); it != nil; it = it.Next() {
		// 有 NPC 的 cell 也不能放置物品
		if _, ok := it.Value.(*npc); ok {
			return true
		}
		if _, ok := it.Value.(*item); ok {
			return true
		}
	}
	return false
}

func (c *cell) isValid() bool {
	return c.attribute == cm.CellAttributeWalk
}

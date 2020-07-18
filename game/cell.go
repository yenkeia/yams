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

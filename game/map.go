package game

import (
	"container/list"
	"fmt"
	"time"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
)

type mirMap struct {
	now        time.Time
	width      int
	height     int
	version    int
	info       *orm.MapInfo
	cells      []*cell
	aoi        *aoiManager
	actionList *actionList
}

func newMirMap(width, height, version int) *mirMap {
	return &mirMap{
		width:      width,
		height:     height,
		version:    version,
		cells:      make([]*cell, width*height),
		aoi:        newAOIManager(width, height),
		actionList: newActionList(),
	}
}

func (mp *mirMap) setCellAttribute(x, y int, attr cm.CellAttribute) {
	c := new(cell)
	c.attribute = attr
	if attr == cm.CellAttributeWalk {
		c.objects = list.New()
	}
	mp.cells[x+y*mp.width] = c
}

func (mp *mirMap) update(now time.Time) {
	mp.now = now
	mp.actionList.execute(now)
}

func (mp *mirMap) broadcast(pos cm.Point, msg interface{}, excludeID int) {
	aoiGrids := mp.aoi.getSurroundGridsByPoint(pos)
	for _, g := range aoiGrids {
		objs := env.getMapObjects(g.getObjectIDs())
		for _, o := range objs {
			if p, ok := env.players[o.getObjectID()]; ok {
				if p.objectID == excludeID {
					continue
				}
				p.enqueue(msg)
			}
		}
	}
}

func (mp *mirMap) inMap(pos cm.Point) bool {
	x := int(pos.X)
	y := int(pos.Y)
	return x >= 0 && x < mp.width && y >= 0 && y < mp.height
}

func (mp *mirMap) getCell(pos cm.Point) *cell {
	x := int(pos.X)
	y := int(pos.Y)
	if mp.inMap(pos) {
		return mp.cells[x+y*mp.width]
	}
	return nil
}

func (mp *mirMap) addObject(obj mapObject) (err error) {
	pos := obj.getPosition()
	c := mp.getCell(pos)
	if c == nil {
		return fmt.Errorf("pos: %s is not walkable", obj.getPosition())
	}
	switch obj := obj.(type) {
	case *player:
		env.players[obj.objectID] = obj
	case *monster:
		env.monsters[obj.objectID] = obj
	case *npc:
		env.npcs[obj.objectID] = obj
	case *item:
		env.items[obj.objectID] = obj
	case *spell:
		env.spells[obj.objectID] = obj
	}
	c.addObject(obj)
	mp.aoi.addObject(obj)
	return
}

func (mp *mirMap) deleteObject(obj mapObject) (err error) {
	pos := obj.getPosition()
	c := mp.getCell(pos)
	if c == nil {
		return fmt.Errorf("pos: %s is not walkable", obj.getPosition())
	}
	c.deleteObject(obj)
	mp.aoi.deleteObject(obj)
	switch obj := obj.(type) {
	case *player:
		delete(env.players, obj.objectID)
	case *monster:
		delete(env.monsters, obj.objectID)
	case *npc:
		delete(env.npcs, obj.objectID)
	case *item:
		delete(env.items, obj.objectID)
	case *spell:
		delete(env.spells, obj.objectID)
	default:
		panic("deleteObject failed.")
	}
	return
}

// 更新对象在地图中的位置
func (mp *mirMap) updateObject(obj mapObject, pos2 cm.Point) (err error) {
	pos1 := obj.getPosition()
	c1 := mp.getCell(pos1)
	c2 := mp.getCell(pos2)
	c1.deleteObject(obj)
	c2.addObject(obj)

	// 更新在 aoi 中的位置
	g1 := mp.aoi.getGridByPoint(pos1)
	g2 := mp.aoi.getGridByPoint(pos2)
	if mp.aoi.updateObject(obj, g1, g2) {
		if o, ok := env.players[obj.getObjectID()]; ok {
			o.enqueueAreaObjects(g1, g2)
		}
	}
	return
}

// TODO
func (mp *mirMap) canSpawnMonster(pos cm.Point) bool {
	// 判断是否 cell walkable
	// 判断是否已经有 player npc monster
	return true
}

// 从p点开始（包含P），由内至外向周围遍历cell。回调函数返回false，停止遍历
func (mp *mirMap) rangeCell(p cm.Point, depth int, fun func(c *cell, x, y int) bool) {
	px, py := int(p.X), int(p.Y)
	for d := 0; d <= depth; d++ {
		for y := py - d; y <= py+d; y++ {
			if y < 0 {
				continue
			}
			if y >= mp.height {
				break
			}
			for x := px - d; x <= px+d; {
				if x >= mp.width {
					break
				}
				if x >= 0 {
					if !fun(mp.getCell(cm.NewPoint(x, y)), x, y) {
						return
					}
				}
				if y-py == d || y-py == -d {
					x++ // x += 1
				} else {
					x += d * 2
				}
			}
		}
	}
}

func (mp *mirMap) rangeObject(p cm.Point, depth int, fun func(mapObject) bool) {
	mp.rangeCell(p, depth, func(c *cell, _, _ int) bool {
		if c != nil && c.objects != nil {
			for it := c.objects.Front(); it != nil; it = it.Next() {
				if !fun(it.Value.(mapObject)) {
					return false
				}
			}
		}
		return true
	})
}

func (mp *mirMap) canWalk(point cm.Point) bool {
	c := mp.getCell(point)
	if !c.isValid() {
		return false
	}
	for it := c.objects.Front(); it != nil; it = it.Next() {
		o := it.Value.(mapObject)
		if o.isBlocking() {
			return false
		}
	}
	return true
}

func (mp *mirMap) validPoint(p cm.Point) bool {
	c := mp.getCell(p)
	if c == nil {
		return false
	}
	return c.isValid()
}

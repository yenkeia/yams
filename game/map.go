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

func (m *mirMap) setCellAttribute(x, y int, attr cm.CellAttribute) {
	c := new(cell)
	c.attribute = attr
	if attr == cm.CellAttributeWalk {
		c.objects = list.New()
	}
	m.cells[x+y*m.width] = c
}

func (m *mirMap) update(now time.Time) {
	m.now = now
	m.actionList.execute(now)
}

func (m *mirMap) broadcast(pos cm.Point, msg interface{}, excludeID int) {
	aoiGrids := m.aoi.getSurroundGridsByPoint(pos)
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

func (m *mirMap) inMap(pos cm.Point) bool {
	x := int(pos.X)
	y := int(pos.Y)
	return x >= 0 && x < m.width && y >= 0 && y < m.height
}

func (m *mirMap) getCell(pos cm.Point) *cell {
	x := int(pos.X)
	y := int(pos.Y)
	if m.inMap(pos) {
		return m.cells[x+y*m.width]
	}
	return nil
}

func (m *mirMap) addObject(obj mapObject) (err error) {
	pos := obj.getPosition()
	c := m.getCell(pos)
	if c == nil {
		return fmt.Errorf("pos: %s is not walkable", obj.getPosition())
	}
	c.addObject(obj)
	m.aoi.addObject(obj)
	return
}

func (m *mirMap) deleteObject(obj mapObject) (err error) {
	pos := obj.getPosition()
	c := m.getCell(pos)
	if c == nil {
		return fmt.Errorf("pos: %s is not walkable", obj.getPosition())
	}
	c.deleteObject(obj)
	m.aoi.deleteObject(obj)
	switch obj := obj.(type) {
	case *player:
		delete(env.players, obj.objectID)
	case *monster:
		delete(env.monsters, obj.objectID)
	case *item:
		delete(env.items, obj.objectID)
	default:
		panic("deleteObject failed.")
	}
	return
}

// 更新对象在地图中的位置
func (m *mirMap) updateObject(obj mapObject, pos2 cm.Point) (err error) {
	pos1 := obj.getPosition()
	c1 := m.getCell(pos1)
	c2 := m.getCell(pos2)
	c1.deleteObject(obj)
	c2.addObject(obj)

	// 更新在 aoi 中的位置
	g1 := m.aoi.getGridByPoint(pos1)
	g2 := m.aoi.getGridByPoint(pos2)
	if g1.gID == g2.gID {
		return
	}
	switch o := obj.(type) {
	case *player:
		o.enqueueAreaObjects(g1, g2)
	case *monster:
		o.broadcastInfo()
	}
	return
}

// TODO
func (m *mirMap) canSpawnMonster(pos cm.Point) bool {
	// 判断是否 cell walkable
	// 判断是否已经有 player npc monster
	return true
}

// 从p点开始（包含P），由内至外向周围遍历cell。回调函数返回false，停止遍历
func (m *mirMap) rangeCell(p cm.Point, depth int, fun func(c *cell, x, y int) bool) {
	px, py := int(p.X), int(p.Y)
	for d := 0; d <= depth; d++ {
		for y := py - d; y <= py+d; y++ {
			if y < 0 {
				continue
			}
			if y >= m.height {
				break
			}
			for x := px - d; x <= px+d; {
				if x >= m.width {
					break
				}
				if x >= 0 {
					if !fun(m.getCell(cm.NewPoint(x, y)), x, y) {
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

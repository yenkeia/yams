package game

import (
	"container/list"
	"fmt"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

type mapObject interface {
	getObjectID() int
	getPosition() cm.Point
}

type mirMap struct {
	width   int
	height  int
	version int
	info    *orm.MapInfo
	cells   []*cell
	aoi     *aoiManager
}

func newMirMap(width, height, version int) *mirMap {
	return &mirMap{
		width:   width,
		height:  height,
		version: version,
		cells:   make([]*cell, width*height),
		aoi:     newAOIManager(width, height),
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

func (m *mirMap) update() {

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
	grids1 := m.aoi.getSurroundGridsByGid(g1.gID)
	grids2 := m.aoi.getSurroundGridsByGid(g2.gID)
	same := make([]*aoiGrid, 0)
	for i := range grids1 {
		for j := range grids2 {
			if grids1[i].gID == grids2[j].gID {
				same = append(same, grids1[i])
			}
		}
	}
	// grids1 - same
	oldGrids := make([]*aoiGrid, 0)
	for i := range grids1 {
		for j := range same {
			if grids1[i].gID == same[j].gID {
				continue
			}
			oldGrids = append(oldGrids, grids1[i])
		}
	}
	// grids2 - same
	newGrids := make([]*aoiGrid, 0)
	for i := range grids2 {
		for j := range same {
			if grids2[i].gID == same[j].gID {
				continue
			}
			newGrids = append(newGrids, grids2[i])
		}
	}
	for i := range oldGrids {
		ids := oldGrids[i].getObjectIDs()
		for _, id := range ids {
			if p, ok := env.players[id]; ok {
				switch obj.(type) {
				case *player:
					p.enqueue(&server.ObjectRemove{ObjectID: uint32(obj.getObjectID())})
				}
			}
		}
	}
	for i := range newGrids {
		ids := newGrids[i].getObjectIDs()
		for _, id := range ids {
			if p, ok := env.players[id]; ok {
				switch o := obj.(type) {
				case *player:
					p.enqueueObjectPlayer(o)
				}
			}
		}
	}
	return
}

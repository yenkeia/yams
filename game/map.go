package game

import (
	"container/list"
	"fmt"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
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

// TODO 更新对象在地图中的位置
func (m *mirMap) updateObject(obj mapObject, pos cm.Point) (err error) {
	return
}

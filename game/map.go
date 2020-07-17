package game

import (
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
)

type mirMap struct {
	width   int
	height  int
	version int
	info    *orm.MapInfo
	cells   []*cell
}

func newMirMap(width, height, version int) *mirMap {
	return &mirMap{
		width:   width,
		height:  height,
		version: version,
		cells:   make([]*cell, width*height),
	}
}

func (m *mirMap) setCellAttribute(x, y int, attr cm.CellAttribute) {
	c := new(cell)
	c.attribute = attr
	if attr == cm.CellAttributeWalk {
		c.objects = make([]interface{}, 0)
	}
	m.cells[x+y*m.width] = c
}

func (m *mirMap) update() {

}

func (m *mirMap) addObject(obj interface{}) (err error) {
	return
}

func (m *mirMap) deleteObject(obj interface{}) (err error) {
	return
}

func (m *mirMap) updateObject(obj interface{}, pos cm.Point) (err error) {
	return
}

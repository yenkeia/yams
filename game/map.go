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
		c.objects = make([]mapObject, 0)
	}
	m.cells[x+y*m.width] = c
}

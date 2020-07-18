package game

import (
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
)

type npc struct {
	name     string
	objectID int
	info     *orm.NPCInfo
}

func newNPC(ni *orm.NPCInfo) *npc {
	n := new(npc)
	n.name = ni.Name
	n.info = ni
	return n
}

func (n *npc) getObjectID() int {
	return n.objectID
}

func (n *npc) getPosition() cm.Point {
	return cm.Point{X: uint32(n.info.LocationX), Y: uint32(n.info.LocationY)}
}

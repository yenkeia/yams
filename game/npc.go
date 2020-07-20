package game

import (
	"time"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

type npc struct {
	info      *orm.NPCInfo
	objectID  int
	name      string
	turnTime  time.Time
	direction cm.MirDirection
	location  cm.Point
}

func newNPC(ni *orm.NPCInfo) *npc {
	n := new(npc)
	n.name = ni.Name
	n.info = ni
	n.direction = cm.MirDirection(cm.RandomInt(0, 1))
	n.location = cm.NewPoint(n.info.LocationX, n.info.LocationY)
	n.turnTime = time.Now()
	return n
}

func (n *npc) getObjectID() int {
	return n.objectID
}

func (n *npc) getPosition() cm.Point {
	return cm.Point{X: uint32(n.info.LocationX), Y: uint32(n.info.LocationY)}
}

func (n *npc) update(now time.Time) {
	if now.After(n.turnTime) {
		n.turnTime = now.Add(time.Second * time.Duration(cm.RandomInt(20, 60)))
		n.direction = cm.MirDirection(cm.RandomInt(0, 1))
		n.broadcast(&server.ObjectTurn{ObjectID: uint32(n.objectID), Location: n.location, Direction: n.direction})
	}
}

func (n *npc) broadcast(msg interface{}) {
	mp := env.maps[n.info.MapID]
	mp.broadcast(n.location, msg)
}

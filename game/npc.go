package game

import (
	"time"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

type npc struct {
	baseObject
	info     *orm.NPCInfo
	turnTime time.Time
	script   *npcScript
}

func newNPC(info *orm.NPCInfo) *npc {
	n := new(npc)
	n.objectID = env.newObjectID()
	n.name = info.Name
	n.nameColor = cm.ColorWhite
	n.mapID = info.MapID
	n.location = cm.NewPoint(info.LocationX, info.LocationY)
	n.direction = cm.MirDirection(cm.RandomInt(0, 1))
	n.info = info
	n.turnTime = time.Now()
	n.script = newNPCScript(conf.Assets + "/NPCs/" + n.info.Filename)
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
	mp.broadcast(n.location, msg, n.objectID)
}

package game

import (
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

type monster struct {
	info       *orm.MonsterInfo
	objectID   int
	mapID      int
	location   cm.Point
	direction  cm.MirDirection
	isDead     bool
	isSkeleton bool
	poison     cm.PoisonType
	isHidden   bool
}

func newMonster(mapID int, location cm.Point, info *orm.MonsterInfo) *monster {
	return &monster{
		info:       info,
		objectID:   env.newObjectID(),
		mapID:      mapID,
		location:   location,
		direction:  cm.RandomDirection(),
		isDead:     false,
		isSkeleton: false,
		poison:     cm.PoisonTypeNone,
		isHidden:   false,
	}
}

func (m *monster) getObjectID() int {
	return m.objectID
}

func (m *monster) getPosition() cm.Point {
	return m.location
}

func (m *monster) broadcast(msg interface{}) {
	mp := env.maps[m.mapID]
	aoiGrids := mp.aoi.getSurroundGridsByPoint(m.location)
	for _, g := range aoiGrids {
		objs := env.getMapObjects(g.getObjectIDs())
		for _, o := range objs {
			if p, ok := env.players[o.getObjectID()]; ok {
				p.enqueue(msg)
			}
		}
	}
}

func (m *monster) broadcastInfo() {
	m.broadcast(&server.ObjectMonster{
		ObjectID:          uint32(m.objectID),
		Name:              m.info.Name,
		NameColor:         cm.ColorWhite.ToInt32(),
		Location:          m.location,
		Image:             cm.Monster(m.info.Image),
		Direction:         m.direction,
		Effect:            uint8(m.info.Effect),
		AI:                uint8(m.info.AI),
		Light:             uint8(m.info.Light),
		Dead:              m.isDead,
		Skeleton:          m.isSkeleton,
		Poison:            m.poison,
		Hidden:            m.isHidden,
		ShockTime:         0,     // TODO
		BindingShotCenter: false, // TODO
		Extra:             false, // TODO
		ExtraByte:         0,     // TODO
	})
}

// TODO
func (m *monster) broadcastHealthChange() {

}

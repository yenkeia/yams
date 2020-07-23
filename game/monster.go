package game

import (
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

type monster struct {
	baseObject
	info       *orm.MonsterInfo
	isDead     bool
	isSkeleton bool
	poison     cm.PoisonType
	isHidden   bool
}

func newMonster(mapID int, location cm.Point, info *orm.MonsterInfo) *monster {
	m := &monster{
		info:       info,
		isDead:     false,
		isSkeleton: false,
		poison:     cm.PoisonTypeNone,
		isHidden:   false,
	}
	m.objectID = env.newObjectID()
	m.name = info.Name
	m.nameColor = cm.ColorWhite
	m.mapID = mapID
	m.location = location
	m.direction = cm.RandomDirection()
	return m
}

func (m *monster) getObjectID() int {
	return m.objectID
}

func (m *monster) getPosition() cm.Point {
	return m.location
}

func (m *monster) broadcast(msg interface{}) {
	mp := env.maps[m.mapID]
	mp.broadcast(m.location, msg, m.objectID)
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

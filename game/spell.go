package game

import (
	"time"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/proto/server"
)

// spell 是地图上显示的魔法 比如火墙
type spell struct {
	base
	spell        cm.Spell
	expireTime   time.Time     // 消失时间
	tickDuration time.Duration // 两次生效时间间隔
	playerID     int           // objectID
	value        int           // 总数值，一般是伤害
}

func newSpell(playerID int, sp cm.Spell, value int,
	mapID int, loc cm.Point,
	tickDuration time.Duration, expireTime time.Time) *spell {
	res := new(spell)
	res.objectID = env.newObjectID()
	res.name = gdb.spellMagicInfoMap[sp].Name
	res.nameColor = cm.ColorWhite
	res.mapID = mapID
	res.location = loc
	res.spell = sp
	res.expireTime = expireTime
	res.tickDuration = tickDuration
	res.playerID = playerID
	res.value = value
	return res
}

func (s *spell) getObjectID() int {
	return s.objectID
}

func (s *spell) getPosition() cm.Point {
	return s.location
}

func (s *spell) isBlocking() bool {
	return false
}

func (s *spell) update(now time.Time) {
	mp := env.maps[s.mapID]
	if now.After(s.expireTime) {
		s.broadcast(&server.ObjectRemove{ObjectID: uint32(s.objectID)})
		mp.deleteObject(s)
		return
	}
	// TODO
	// obj.attacked
}

func (s *spell) broadcast(msg interface{}) {
	env.maps[s.mapID].broadcast(s.location, msg, s.objectID)
}

// TODO
func (s *spell) broadcastInfo() {
	switch s.spell {
	case cm.SpellHealing:
	default:
		s.broadcast(&server.ObjectSpell{
			ObjectID:  uint32(s.objectID),
			Location:  s.location,
			Spell:     s.spell,
			Direction: s.direction,
		})
	}
}

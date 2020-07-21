package game

import (
	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
)

type respawn struct {
	info *orm.RespawnInfo
}

func newRespawn(ri *orm.RespawnInfo) *respawn {
	return &respawn{info: ri}
}

func (r *respawn) spawn() {
	mp := env.maps[r.info.MapID]
	for i := 0; i < 10; i++ {
		x := r.info.LocationX + cm.RandomInt(-r.info.Spread, r.info.Spread)
		y := r.info.LocationY + cm.RandomInt(-r.info.Spread, r.info.Spread)
		if !mp.canSpawnMonster(cm.NewPoint(x, y)) {
			continue
		}
		m := newMonster(mp.info.ID, cm.NewPoint(x, y), gdb.monsterInfoMap[r.info.MonsterID])

		// TODO 从 env 中删除
		env.monsters[m.objectID] = m
		mp.addObject(m)

		m.broadcastInfo()
		m.broadcastHealthChange()
		return
	}
}

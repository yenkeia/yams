package game

import "github.com/yenkeia/yams/game/orm"

type npc struct {
	name string
	info *orm.NPCInfo
}

func newNPC(ni *orm.NPCInfo) *npc {
	n := new(npc)
	n.name = ni.Name
	n.info = ni
	return n
}

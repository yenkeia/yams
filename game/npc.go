package game

import (
	"time"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

type npc struct {
	base
	info     *orm.NPCInfo
	turnTime time.Time
	script   *npcScript
	goods    []*userItem
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
	n.loadGoods()
	return n
}

func (n *npc) getObjectID() int {
	return n.objectID
}

func (n *npc) getPosition() cm.Point {
	return cm.Point{X: uint32(n.info.LocationX), Y: uint32(n.info.LocationY)}
}

func (n *npc) isBlocking() bool {
	return true
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

func (n *npc) loadGoods() {
	n.goods = make([]*userItem, 0)
	for _, name := range n.script.trade {
		n.goods = append(n.goods, newUserItem(gdb.itemInfoNameMap[name]))
	}
}

func (n *npc) processSpecial(p *player, key string) {
	switch key {
	case BuyKey:
	case SellKey:
	case BuySellKey:
		ls := make([]*server.UserItem, 0)
		for _, good := range n.goods {
			p.enqueueItemInfo(good.info.ID)
			ls = append(ls, good.serverUserItem())
		}
		p.enqueue(&server.NPCGoods{Goods: ls, Rate: 1.0, Type: cm.PanelTypeBuy})
	case RepairKey:
	case SRepairKey:
	case CraftKey:
	case RefineKey:
	case RefineCheckKey:
	case RefineCollectKey:
	case ReplaceWedRingKey:
	case StorageKey:
	case BuyBackKey:
	case BuyUsedKey:
	case ConsignKey:
	case MarketKey:
	case ConsignmentsKey:
	case GuildCreateKey:
	case RequestWarKey:
	case SendParcelKey:
	case CollectParcelKey:
	case AwakeningKey:
	case DisassembleKey:
	case DowngradeKey:
	case ResetKey:
	case PearlBuyKey:
	}
}

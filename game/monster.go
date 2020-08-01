package game

import (
	"fmt"
	"time"

	"github.com/yenkeia/yams/game/cm"
	"github.com/yenkeia/yams/game/orm"
	"github.com/yenkeia/yams/game/proto/server"
)

type monster struct {
	baseObject
	respawnID    int
	info         *orm.MonsterInfo
	isDead       bool
	isSkeleton   bool
	poison       cm.PoisonType
	isHidden     bool
	hp           int
	mp           int
	level        int
	petLevel     int
	experience   int
	maxHP        int
	minAC        int
	maxAC        int
	minMAC       int
	maxMAC       int
	minDC        int
	maxDC        int
	minMC        int
	maxMC        int
	minSC        int
	maxSC        int
	expOwnerID   int // 获得经验的玩家 objectID
	expOwnerTime time.Time
	masterID     int           // 怪物主人 objectID
	deleteTime   time.Time     // 从 env.monsters 中删除的时间
	bt           *behaviorTree // 怪物行为树
}

func newMonster(respawnID int, mapID int, location cm.Point, info *orm.MonsterInfo) *monster {
	m := &monster{
		respawnID:    respawnID,
		info:         info,
		isDead:       false,
		isSkeleton:   false,
		poison:       cm.PoisonTypeNone,
		isHidden:     false,
		hp:           info.HP,
		maxHP:        info.HP,
		level:        info.Level,
		experience:   info.Experience,
		minAC:        info.MinAC,
		maxAC:        info.MaxAC,
		minMAC:       info.MinMAC,
		maxMAC:       info.MaxMAC,
		minDC:        info.MinDC,
		maxDC:        info.MaxDC,
		minMC:        info.MinMC,
		maxMC:        info.MaxMC,
		minSC:        info.MinSC,
		maxSC:        info.MaxSC,
		expOwnerID:   0,
		expOwnerTime: time.Now(),
		masterID:     0,
	}
	m.objectID = env.newObjectID()
	m.name = info.Name
	m.nameColor = cm.ColorWhite
	m.mapID = mapID
	m.location = location
	m.direction = cm.RandomDirection()
	m.bt = newBehaviorTree(m)
	return m
}

func (m *monster) String() string {
	return fmt.Sprintf("怪物: %s", m.name)
}

func (m *monster) getObjectID() int {
	return m.objectID
}

func (m *monster) getPosition() cm.Point {
	return m.location
}

// 怪物定时轮询
func (m *monster) update(now time.Time) {
	if m.isDead && now.After(m.deleteTime) {
		log.Debugf("怪物[%s,%d]死亡，从游戏环境删除", m.name, m.objectID)
		m.broadcast(&server.ObjectRemove{ObjectID: uint32(m.objectID)})
		env.maps[m.mapID].deleteObject(m)
		return
	}
	if m.expOwnerID != 0 && now.After(m.expOwnerTime) {
		m.expOwnerID = 0
		log.Debugln("monster expOwnerID = 0")
	}
	m.bt.process(now)
}

// ChangeHP 怪物改变血量 amount 可以是负数(扣血)
func (m *monster) changeHP(amount int) {
	if m.isDead {
		return
	}
	log.Debugf("monster changeHP. 当前血量 m.hp: %d, 变化量 amount: %d.", m.hp, amount)
	value := m.hp + amount
	if value == m.hp {
		return
	}
	if value <= 0 {
		m.die()
		m.hp = 0
	} else {
		m.hp = value
	}
	percent := uint8(float32(m.hp) / float32(m.maxHP) * 100)
	log.Debugf("怪物最终血量 m.hp: %d, m.maxHP: %d, percent: %d\n", m.hp, m.maxHP, percent)
	m.broadcast(&server.ObjectHealth{
		ObjectID: uint32(m.objectID),
		Percent:  percent,
		Expire:   5,
	})
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

func (m *monster) broadcastHealthChange() {
	percent := byte(float32(m.hp) / float32(m.maxHP) * 100)
	msg := &server.ObjectHealth{
		ObjectID: uint32(m.objectID),
		Percent:  percent,
		Expire:   5,
	}
	m.broadcast(msg)
}

func (m *monster) broadcastObjectStruck(a attacker) {
	attackerID := 0
	switch atk := a.(type) {
	case *player:
		attackerID = atk.objectID
	case *monster:
		attackerID = atk.objectID
	}
	m.broadcast(&server.ObjectStruck{
		ObjectID:   uint32(m.objectID),
		AttackerID: uint32(attackerID),
		LocationX:  int32(m.location.X),
		LocationY:  int32(m.location.Y),
		Direction:  m.direction,
	})
}

func (m *monster) broadcastDamageIndicator(typ cm.DamageType, dmg int) {
	m.broadcast(&server.DamageIndicator{Damage: int32(dmg), Type: typ, ObjectID: uint32(m.objectID)})
}

// TODO
func (m *monster) attacked(atk attacker, dmg int, typ cm.DefenceType, isWeapon bool) int {
	log.Debugf("monster[%s] attacked. attacker: [%s], damage: %d", m, atk, dmg)
	armor := 0    // TODO
	damage := dmg // TODO
	value := damage - armor
	log.Debugf("attacker damage: %d, monster armour: %d\n", damage, armor)
	if value <= 0 {
		m.broadcastDamageIndicator(cm.DamageTypeMiss, 0)
		return 0
	}

	// 判断怪物被谁攻击，设置 expOwner
	switch atk := atk.(type) {
	case *monster:
		if atk.masterID != 0 {
			m.expOwnerID = atk.masterID
		} else {
			m.expOwnerID = 0
		}
	case *player:
		m.expOwnerID = atk.objectID
	}
	if m.expOwnerID != 0 {
		m.expOwnerTime = time.Now().Add(5 * time.Second)
	}
	log.Debugf("monster attacked. expOwnerID: %d, expOwnerTime: %s", m.expOwnerID, m.expOwnerTime)

	// TODO 还有很多没做
	m.broadcastObjectStruck(atk)
	m.broadcastDamageIndicator(cm.DamageTypeHit, -value)
	m.changeHP(-value)
	return 0
}

// TODO
func (m *monster) isAttackTarget(attacker) bool {
	return true
}

// TODO
func (m *monster) attack(...interface{}) {

}

func (m *monster) die() {
	if m.isDead {
		return
	}
	m.hp = 0
	m.isDead = true
	m.broadcast(&server.ObjectDied{
		ObjectID:  uint32(m.objectID),
		LocationX: int32(m.location.X),
		LocationY: int32(m.location.Y),
		Direction: m.direction,
		Type:      0,
	})
	m.drop()
	// 击杀者获得经验
	if m.expOwnerID != 0 && m.masterID == 0 {
		p, ok := env.players[m.expOwnerID]
		if !ok {
			return
		}
		log.Debugf("怪物[%s]死亡。击杀者[%s]", m.name, p.name)
		p.winExp(m.experience, m.level)
	}

	// 设置怪物从 env.monsters 中删除的时间，在 monster.update 时候再删除
	mp := env.maps[m.mapID]
	m.deleteTime = mp.now.Add(10 * time.Second)
	// 往地图中加入一个延迟动作，刷一个新的怪物
	r := env.respawns[m.respawnID]
	mp.actionList.pushDelayAction(cm.DelayedTypeSpawn, r.info.Interval*1000, func() {
		r.spawnOneMonster()
	})
}

// TODO 怪物掉落
func (m *monster) drop() {

}

package game

import (
	"time"

	"github.com/yenkeia/yams/game/cm"
)

type base struct {
	objectID  int
	name      string
	nameColor cm.Color
	mapID     int
	location  cm.Point
	direction cm.MirDirection
}

type mapObject interface {
	getObjectID() int
	getPosition() cm.Point
	isBlocking() bool // 玩家是否能穿过
	update(time.Time)
}

// 攻击者
type attacker interface {
	attack(...interface{})
}

// 可以被攻击的对象
type attackTarget interface {
	// Attacked(attacker, damage, DefenceType , damageWeapon );
	mapObject
	attacked(attacker, int, cm.DefenceType, bool) int
	isAttackTarget(attacker) bool   // 目标是攻击者的队友或宠物
	isFriendlyTarget(attacker) bool // 目标是攻击者的攻击对象
	changeHP(int)                   // 改变血量，如果传入负数表示扣血
	applyPoison(*poison, attacker)
}

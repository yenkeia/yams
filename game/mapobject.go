package game

import (
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
	isBlocking() bool
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
	isAttackTarget(attacker) bool
}

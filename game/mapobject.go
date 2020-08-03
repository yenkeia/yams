package game

import "github.com/yenkeia/yams/game/cm"

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
}

// 攻击者
type attacker interface {
	attack(...interface{})
	getAttackTarget(int) attackTarget
}

// 可以被攻击的对象
type attackTarget interface {
	// Attacked(attacker, damage, DefenceType , damageWeapon );
	getObjectID() int
	attacked(attacker, int, cm.DefenceType, bool) int
	isAttackTarget(attacker) bool
}

package game

import "github.com/yenkeia/yams/game/cm"

type baseObject struct {
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
}

// 可以被攻击的对象
type attackableObject interface {
	// Attacked(attacker, damage, DefenceType , damageWeapon );
	attacked(attacker, int, cm.DefenceType, bool) int
	isAttackTarget(attacker) bool
}

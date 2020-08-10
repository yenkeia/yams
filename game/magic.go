package game

import (
	"fmt"
	"time"

	"github.com/yenkeia/yams/game/cm"
)

type magicContext struct {
	spell       cm.Spell
	target      attackTarget
	player      *player
	targetPoint cm.Point
}

func startMagic(ctx *magicContext) (targetID int, err error) {
	if ctx.target != nil {
		targetID = 0
		if ctx.target.isAttackTarget(ctx.player) {
			targetID = ctx.target.getObjectID()
		}
	}
	switch ctx.spell {
	case cm.SpellFireBall:
		fireBall(ctx)
	default:
		return 0, fmt.Errorf("技能暂未实现")
	}
	return
}

func fireBall(ctx *magicContext) {
	player := ctx.player
	target := ctx.target
	magic := ctx.player.magics[ctx.spell]
	if target == nil || !target.isAttackTarget(ctx.player) {
		return
	}
	damage := magic.getDamage(player.getAttackPower(player.minMC, player.maxMC))
	delay := cm.MaxDistance(player.location, target.getPosition())*50 + 500
	player.actionList.pushDelayAction(cm.DelayedTypeMagic, time.Duration(delay)*time.Millisecond, func() {
		target.attacked(player, damage, cm.DefenceTypeMAC, false)
	})
}

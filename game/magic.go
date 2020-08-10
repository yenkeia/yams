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
	case cm.SpellFireWall:
		fireWall(ctx)
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

func fireWall(ctx *magicContext) {
	player := ctx.player
	magic := ctx.player.magics[ctx.spell]
	location := ctx.targetPoint
	value := magic.getDamage(player.getAttackPower(player.minMC, player.maxMC))
	mp := env.maps[player.mapID]
	mp.actionList.pushDelayAction(cm.DelayedTypeMagic, time.Millisecond*time.Duration(500), func() {
		dir := cm.MirDirectionUp
		points := []cm.Point{}
		for i := 0; i < 4; i++ {
			point := location.NextPoint(dir, 1)
			dir = dir.NextDirection().NextDirection()
			points = append(points, point)
		}
		points = append(points, location)
		for _, point := range points {
			if mp.validPoint(point) {
				c := mp.getCell(point)
				cast := true
				for it := c.objects.Front(); it != nil; it = it.Next() {
					if obj, ok := it.Value.(*spell); ok {
						if obj.spell == cm.SpellFireWall {
							cast = false
							break
						}
					}
				}
				if cast {
					expireTime := mp.now.Add(time.Millisecond * time.Duration((10+value/2)*1000))
					tickSpeed := time.Duration(2000 * time.Millisecond)
					s := newSpell(player.objectID, ctx.spell, value, mp.info.ID, point, tickSpeed, expireTime)
					mp.addObject(s)
					s.broadcastInfo()
				}
			}
		}
	})
}
